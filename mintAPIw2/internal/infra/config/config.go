package config

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/log"
	"github.com/spf13/viper"
)

const (
	PortConfig      = "server.port"
	NameConfig      = "server.name"
	EthQuickURL     = "ethereum.quicknode"
	ContractAdress  = "ethereum.contract_address"
	OwnerPrivateKey = "owner.privatekey"
	QRDir           = "qr.dir"
	IPFSURI         = "ipfs.url"
)

const (
	portConfigDefaultValue = "3000"
	NameConfigDefaultValue = "meent"
)

// Configurations are configuration settings for the application
type GlobalConfigurations struct {
	ConfigurationAccess *viper.Viper
}

var Configurations *GlobalConfigurations

// Section to handle global conf on infrastructure
func LoadConfiguration() *viper.Viper {
	configs := viper.New()

	workingdir, err := os.Getwd()
	if err != nil {
		log.Error(err.Error())
	}

	configs.SetConfigName("config")
	configs.SetConfigType("toml")
	configs.AddConfigPath(workingdir)

	err = configs.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	configs.SetDefault(NameConfig, NameConfigDefaultValue)
	configs.SetDefault(PortConfig, portConfigDefaultValue)

	Configurations = &GlobalConfigurations{
		ConfigurationAccess: configs,
	}

	return configs
}

func GetGlobalConfigs() *GlobalConfigurations {
	return Configurations
}

func GetPortConfigs() string {
	return Configurations.ConfigurationAccess.GetString(PortConfig)
}

func GetNameConfigs() string {
	return Configurations.ConfigurationAccess.GetString(NameConfig)
}

func GetQuickNodeConfigs() string {
	return Configurations.ConfigurationAccess.GetString(EthQuickURL)
}

func GetContractAdressConfigs() string {
	return Configurations.ConfigurationAccess.GetString(ContractAdress)
}

func GetOwnerPrivateKey() string {
	return Configurations.ConfigurationAccess.GetString(OwnerPrivateKey)
}

func GetQRDirectory() string {
	return Configurations.ConfigurationAccess.GetString(QRDir)
}

func GetIPFSURL() string {
	return Configurations.ConfigurationAccess.GetString(IPFSURI)
}
