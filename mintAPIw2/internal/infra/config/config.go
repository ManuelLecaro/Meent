package config

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/log"
	"github.com/spf13/viper"
)

const (
	PortConfig                = "server.port"
	portConfigDefaultValue    = "3000"
	NameConfig                = "server.name"
	NameConfigDefaultValue    = "mintAPI"
	VersionConfig             = "VERSION"
	VersionConfigDefaultValue = "0.0.1"
	DebugConfig               = "DEBUG"
	DebugConfigDefaultValue   = false
	OwnerPrivateKey           = "OWNER_PRIVATE_KEY"
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
