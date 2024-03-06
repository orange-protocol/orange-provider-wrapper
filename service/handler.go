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
	Error  *HttpError `json:"error"`
}
type HttpError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
type ErrorCode string

type ResponseDataWithSig struct {
	Data string `json:"data"`
	Sig  string `json:"sig"`
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
		log.Error("GetStatus Marshal", "err", err.Error())
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
	orangeAddr, err := utils.DIDToEthAddress(config.GlobalConfig.OrangeDID)
	if err != nil {
		return false, err
	}
	sigbytes, err := hexutil.Decode(sig)
	if err != nil {
		return false, err
	}
	return utils.ETHVerifySig(orangeAddr, sigbytes, []byte(requestData)), nil
}

func (sp *ProxyService) GetParamFromRequest(r *http.Request) (*OrangeRequest, error) {
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

func (sp *ProxyService) GenerateHandleFunc(cfg config.APIConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Infof("Generating handle:%s...", cfg.ServerPath)
		param, err := sp.GetParamFromRequest(r)
		if err != nil {
			doResponse(w, NewHttpError(INVALID_PARAM, err.Error()), nil)
			return
		}
		//check sig
		valid, err := sp.VerifyRequestSignature(param.Request.RequestData, param.Sig)
		if !valid || err != nil {
			doResponse(w, NewHttpError(INVALID_PARAM, err.Error()), nil)
			return
		}

		var req *http.Request
		switch strings.ToUpper(cfg.ParamType) {
		case "BODY":
			payload := strings.NewReader(param.Request.RequestData)
			req, err = http.NewRequest(cfg.ApiMethod, cfg.ApiUrl, payload)
		case "URL":
			url := fmt.Sprintf("%s/%s", cfg.ApiUrl, param.Request.RequestData)
			req, err = http.NewRequest(cfg.ApiMethod, url, nil)
		default:
			log.Errorf("unsupported paramType: %v", cfg.ParamType)
			err = fmt.Errorf("unsupported paramType: %v", cfg.ParamType)
		}
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
		if res.StatusCode != 200 {
			doResponse(w, NewHttpError(INTERNAL_ERROR, fmt.Sprintf("status code:%d", res.StatusCode)), nil)
			return
		}

		data, err := io.ReadAll(res.Body)
		if err != nil {
			doResponse(w, NewHttpError(INTERNAL_ERROR, fmt.Sprintf("status code:%d", res.StatusCode)), nil)
			return
		}

		sig, err := GlobalSignerService.SignMsg(data)
		if err != nil {
			doResponse(w, NewHttpError(INTERNAL_ERROR, fmt.Sprintf("status code:%d", res.StatusCode)), nil)
			return
		}

		dataWithSig := &ResponseDataWithSig{
			Data: string(data),
			Sig:  hexutil.Encode(sig),
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
