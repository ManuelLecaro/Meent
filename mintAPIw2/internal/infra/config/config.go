package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	portConfig             = "PORT"
	portConfigDefaultValue = "3000"
	NameConfig             = "NAME"
	NameConfigDefaultValue = "mintAPI"
)

// Configurations are configuration settings for the application
type GlobalConfigurations struct {
	ConfigurationAccess *viper.Viper
}

var Configurations *GlobalConfigurations

// Section to handle global conf on infrastructure
func LoadConfiguration() *viper.Viper {
	configs := viper.New()

	configs.SetConfigName("config")
	configs.SetConfigType("toml")
	configs.AddConfigPath("/../../..")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	configs.SetDefault(portConfig, portConfigDefaultValue)
	configs.SetDefault(NameConfig, NameConfigDefaultValue)

	Configurations = &GlobalConfigurations{
		ConfigurationAccess: configs,
	}

	return configs
}
