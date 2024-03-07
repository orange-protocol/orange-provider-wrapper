package service

import (
	"orange-provider-wrapper/config"
	"orange-provider-wrapper/log"
)

func InitAllServices() error {
	InitProxyService()
	err := InitDidService(config.GlobalConfig.ChainRpc, config.GlobalConfig.ContractAddress)
	if err != nil {
		log.Errorf("InitDidService failed: %v", err)
		return err
	}
	err = InitSignerService(config.GlobalConfig.Keystore, config.GlobalConfig.WalletPwd)
	if err != nil {
		log.Errorf("InitSignerService failed: %v", err)
		return err
	}
	return nil
}
