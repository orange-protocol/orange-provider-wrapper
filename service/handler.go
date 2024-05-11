package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"orange-provider-wrapper/config"
	"orange-provider-wrapper/log"
	"orange-provider-wrapper/utils"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

var GlobalProxyService *ProxyService

type ProxyService struct {
	client *http.Client
}

type OrangeRequest struct {
	Request struct {
		RequestDid  string `json:"requester_did"`
		RequestData string `json:"request_data"`
		Encrypt     bool   `json:"encrypt"`
	} `json:"request"`
	Sig string `json:"sig"`
}

type HttpResult struct {
	Result any        `json:"result"`
	Error  *HttpError `json:"error,omitempty"`
}
type HttpError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
type ErrorCode string

type ResponseDataWithSig struct {
	Data        string `json:"data"`
	ProviderDid string `json:"provider_did"`
	Sig         string `json:"sig"`
}

type ResponseData struct {
	Data      *ResponseDataWithSig `json:"data"`
	Encrypted string               `json:"encrypted"`
}

var (
	INVALID_PARAM  ErrorCode = "INVALID_PARAM"
	INTERNAL_ERROR ErrorCode = "INTERNAL_ERROR"
	NOT_IMPLEMENT  ErrorCode = "NOT_IMPLEMENT"
)

func doResponse(w http.ResponseWriter, herr *HttpError, result any) {
	resp := ToHttpResult(result, herr)
	bts, err := json.Marshal(resp)
	if err != nil {
		log.Errorf("GetStatus Marshal", "err", err.Error())
		http.Error(w, http.StatusText(500), 500)
		return
	}
	w.Write(bts)
}
func ToHttpResult(result any, herr *HttpError) HttpResult {
	if herr != nil {
		return HttpResult{Result: result, Error: herr}
	} else {
		return HttpResult{Result: result, Error: nil}
	}
}
func NewHttpError(code ErrorCode, msg string) *HttpError {
	return &HttpError{Code: string(code), Message: msg}
}

func InitProxyService() {
	GlobalProxyService = &ProxyService{client: new(http.Client)}
}

func (sp *ProxyService) VerifyRequestSignature(requestData string, sig string) (bool, error) {
	//
	log.Infof("did: %s", config.GlobalConfig.OrangeDID)
	orangeAddr, err := utils.DIDToEthAddress(config.GlobalConfig.OrangeDID)
	if err != nil {
		log.Errorf("DIDToEthAddress :error: %v", err.Error())
		return false, err
	}
	sigbytes, err := hexutil.Decode(sig)
	if err != nil {
		log.Errorf("Decode :error: %v", err.Error())
		return false, err
	}
	return utils.ETHVerifySig(orangeAddr, sigbytes, []byte(requestData)), nil
}

func (sp *ProxyService) GetDPParamFromRequest(r *http.Request) (*OrangeRequest, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	param := OrangeRequest{}
	err = json.Unmarshal(body, &param)
	if err != nil {
		return nil, err
	}
	return &param, nil
}

func (sp *ProxyService) GetAPParamFromRequest(r *http.Request) (*ResponseData, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	param := ResponseData{}
	err = json.Unmarshal(body, &param)
	if err != nil {
		return nil, err
	}
	return &param, nil
}

func (sp *ProxyService) GenerateDPHandleFunc(cfg config.APIConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		param, err := sp.GetDPParamFromRequest(r)
		if err != nil {
			log.Errorf("GetDPParamFromRequest Error: %v\n", err)

			doResponse(w, NewHttpError(INVALID_PARAM, err.Error()), nil)
			return
		}
		//check sig
		if cfg.VerifyRequest {
			msg, err := json.Marshal(param.Request)
			if err != nil {
				log.Errorf("Marshal Error: %v\n", err)

				doResponse(w, NewHttpError(INVALID_PARAM, err.Error()), nil)
				return
			}
			valid, err := sp.VerifyRequestSignature(string(msg), param.Sig)
			if !valid || err != nil {
				log.Errorf("verify DP signature:data:%s,sig:%s", msg, param.Sig)
				doResponse(w, NewHttpError(INVALID_PARAM, "invalid signature"), nil)
				return
			}
		}
		var req *http.Request
		switch strings.ToUpper(cfg.ParamType) {
		case "BODY":
			payload := strings.NewReader(param.Request.RequestData)
			req, err = http.NewRequest(cfg.ApiMethod, cfg.ApiUrl, payload)
		case "URL":
			url := fmt.Sprintf("%s?%s", cfg.ApiUrl, param.Request.RequestData)
			log.Debugf("URL :%s\n", url)
			req, err = http.NewRequest(cfg.ApiMethod, url, nil)
		case "REST":
			if strings.Contains(cfg.ApiUrl, "$PARAM") {
				p, err := utils.GetRestValueFromParam(param.Request.RequestData)
				if err != nil {
					log.Errorf("GetRestValueFromParam error: %v", err)
					doResponse(w, NewHttpError(INTERNAL_ERROR, err.Error()), nil)
					return
				}
				url := strings.ReplaceAll(cfg.ApiUrl, "$PARAM", p)
				req, err = http.NewRequest(cfg.ApiMethod, url, nil)
			} else {
				req, err = http.NewRequest(cfg.ApiMethod, cfg.ApiUrl, nil)
			}
		default:
			log.Errorf("unsupported paramType: %v", cfg.ParamType)
			err = fmt.Errorf("unsupported paramType: %v", cfg.ParamType)
		}
		if err != nil {
			log.Errorf("send request error: %v", err)
			doResponse(w, NewHttpError(INTERNAL_ERROR, err.Error()), nil)
			return
		}
		if cfg.HasApiKey {
			if strings.EqualFold(cfg.ApiKeyLocation, "HEADER") {
				req.Header.Set(cfg.ApiKeyName, cfg.ApiKey)
			}

			//todo other place to store apikey ???
		}
		if strings.EqualFold(cfg.ParamType, "BODY") {
			req.Header.Add("Content-Type", "application/json")
		}

		res, err := sp.client.Do(req)
		if err != nil {
			doResponse(w, NewHttpError(INTERNAL_ERROR, err.Error()), nil)
			return
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			doResponse(w, NewHttpError(INTERNAL_ERROR, fmt.Sprintf("status code:%d", res.StatusCode)), nil)
			return
		}

		data, err := io.ReadAll(res.Body)
		if err != nil {
			doResponse(w, NewHttpError(INTERNAL_ERROR, fmt.Sprintf("error:%s", err.Error())), nil)
			return
		}
		if sp.checkResponseFailed(data, cfg.FailedKeywords) {
			doResponse(w, NewHttpError(INTERNAL_ERROR, fmt.Sprintf("api returns fail:%s", data)), nil)
			return
		}

		sig, err := GlobalSignerService.SignMsg(data)
		if err != nil {
			doResponse(w, NewHttpError(INTERNAL_ERROR, fmt.Sprintf("status code:%d", res.StatusCode)), nil)
			return
		}

		dataWithSig := &ResponseDataWithSig{
			Data:        string(data),
			ProviderDid: utils.EthAddressToDID(GlobalSignerService.WalletAddress),
			Sig:         hexutil.Encode(sig),
		}

		if param.Request.Encrypt {
			requesterPubkey, err := GlobalDidService.GetDidPublicKey(param.Request.RequestDid)
			if err != nil {
				doResponse(w, NewHttpError(INTERNAL_ERROR, err.Error()), nil)
				return
			}

			if len(requesterPubkey) == 0 {
				doResponse(w, NewHttpError(INTERNAL_ERROR, fmt.Sprintf("requester %s not registered", param.Request.RequestDid)), nil)
				return
			}
			pubkeyBytes, err := hexutil.Decode(requesterPubkey)
			if err != nil {
				doResponse(w, NewHttpError(INTERNAL_ERROR, err.Error()), nil)
				return
			}
			dataToEncrypt, err := json.Marshal(dataWithSig)
			if err != nil {
				doResponse(w, NewHttpError(INTERNAL_ERROR, err.Error()), nil)
				return
			}
			encryptedData, err := utils.EncryptMessageWithPubkey(dataToEncrypt, pubkeyBytes)
			if err != nil {
				doResponse(w, NewHttpError(INTERNAL_ERROR, err.Error()), nil)
				return
			}
			doResponse(w, nil, &ResponseData{
				Data:      nil,
				Encrypted: hexutil.Encode(encryptedData),
			})
			return
		} else {
			doResponse(w, nil, &ResponseData{
				Data:      dataWithSig,
				Encrypted: "",
			})
			return
		}
	}
}

func (sp *ProxyService) GenerateAPHandleFunc(cfg config.APIConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		param, err := sp.GetAPParamFromRequest(r)
		if err != nil {
			log.Errorf("GetAPParamFromRequest Error: %v\n", err)
			doResponse(w, NewHttpError(INVALID_PARAM, err.Error()), nil)
			return
		}
		var dataWithSig *ResponseDataWithSig
		if len(param.Encrypted) > 0 {
			//decrypt param
			data, err := hexutil.Decode(param.Encrypted)
			if err != nil {
				log.Errorf("Decode Error: %v\n", err)
				doResponse(w, NewHttpError(INVALID_PARAM, err.Error()), nil)
				return
			}
			decryptedData, err := utils.DecryptMessage(data, GlobalSignerService.GetPrivateKey())
			if err != nil {
				log.Errorf("DecryptMessage Error: %v\n", err)
				doResponse(w, NewHttpError(INVALID_PARAM, err.Error()), nil)
				return
			}
			dataWithSig = &ResponseDataWithSig{}
			err = json.Unmarshal(decryptedData, dataWithSig)
			if err != nil {
				log.Errorf("Unmarshal Error: %v\n", err)
				doResponse(w, NewHttpError(INVALID_PARAM, err.Error()), nil)
				return
			}
		} else {
			dataWithSig = param.Data
		}

		//verify signature
		if cfg.VerifyRequest {
			verified, err := sp.checkDataWithSig(dataWithSig)
			if err != nil {
				log.Errorf("error checking signature")
				doResponse(w, NewHttpError(INVALID_PARAM, err.Error()), nil)
				return
			}
			if !verified {
				log.Errorf("error checking signature")
				doResponse(w, NewHttpError(INVALID_PARAM, "signature invalid"), nil)
				return
			}
		}

		// send request
		payload := strings.NewReader(dataWithSig.Data)
		//for AP, only POST allowed
		req, err := http.NewRequest("POST", cfg.ApiUrl, payload)
		if err != nil {
			doResponse(w, NewHttpError(INTERNAL_ERROR, err.Error()), nil)
			return
		}
		if cfg.HasApiKey {
			if strings.EqualFold(cfg.ApiKeyLocation, "HEADER") {
				req.Header.Set(cfg.ApiKeyName, cfg.ApiKey)
			}
			//todo other place to store apikey ???
		}
		res, err := sp.client.Do(req)
		if err != nil {
			doResponse(w, NewHttpError(INTERNAL_ERROR, err.Error()), nil)
			return
		}
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			doResponse(w, NewHttpError(INTERNAL_ERROR, fmt.Sprintf("status code:%d", res.StatusCode)), nil)
			return
		}
		data, err := io.ReadAll(res.Body)
		if err != nil {
			doResponse(w, NewHttpError(INTERNAL_ERROR, err.Error()), nil)
			return
		}
		if sp.checkResponseFailed(data, cfg.FailedKeywords) {
			doResponse(w, NewHttpError(INTERNAL_ERROR, fmt.Sprintf("api returns fail:%s", data)), nil)
			return
		}
		sig, err := GlobalSignerService.SignMsg(data)
		if err != nil {
			doResponse(w, NewHttpError(INTERNAL_ERROR, err.Error()), nil)
			return
		}
		resp := &ResponseDataWithSig{
			Data:        string(data),
			ProviderDid: utils.EthAddressToDID(GlobalSignerService.WalletAddress),
			Sig:         hexutil.Encode(sig),
		}
		doResponse(w, nil, resp)
	}
}

func (sp *ProxyService) checkDataWithSig(dws *ResponseDataWithSig) (bool, error) {
	if dws == nil {
		return false, fmt.Errorf("dataWithSig is nil")
	}
	sigbytes, err := hexutil.Decode(dws.Sig)
	if err != nil {
		return false, err
	}
	signer, err := utils.DIDToEthAddress(dws.ProviderDid)
	if err != nil {
		return false, err
	}

	return utils.ETHVerifySig(signer, sigbytes, []byte(dws.Data)), nil
}

func (sp *ProxyService) checkResponseFailed(resp []byte, keywords []string) bool {
	for _, keyword := range keywords {
		if strings.Contains(string(resp), keyword) {
			return true
		}
	}
	return false
}
