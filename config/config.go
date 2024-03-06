package config

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

const (
	ETH_MONITOR_INTERVAL = 3 * time.Second
	ETH_USEFUL_BLOCK_NUM = 6
)

var GlobalConfig *SysConfig

type SysConfig struct {
	OrangeDID       string      `json:"orange_did"`
	WalletFile      string      `json:"wallet_file"`
	WalletPwd       string      `json:"wallet_pwd"`
	Chain           string      `json:"chain"`
	ContractAddress string      `json:"contract_address"`
	APIConfigs      []APIConfig `json:"api_configs"`
}

type APIConfig struct {
	ServerPath     string `json:"server_path"`
	HasApiKey      bool   `json:"has_api_key"`
	ApiKeyLocation string `json:"api_key_location"`
	ApiKeyName     string `json:"api_key_name"`
	ApiKey         string `json:"api_key"`
	ApiUrl         string `json:"api_url"`
	ApiMethod      string `json:"api_method"`
	ParamType      string `json:"param_type"`
}

func LoadConfig(filepath string) error {
	return loadConfigFromFile(filepath)
}

func loadConfigFromFile(filepath string) error {
	configData, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}
	cfg := &SysConfig{}
	err = json.Unmarshal(configData, cfg)
	if err != nil {
		return err
	}
	GlobalConfig = cfg
	return nil
}
