package config

import (
	"bytes"
	"fmt"

	"github.com/spf13/viper"
)

type ViperConfig struct {
	viper *viper.Viper
	env   Environment
}

func (v *ViperConfig) Get(key string, result any) error {
	sub := v.viper.Sub(key)
	if sub != nil {
		return v.viper.Sub(key).Unmarshal(result)
	}
	return fmt.Errorf("empty config entry for: %s", key)
}

func (v *ViperConfig) GetEnv() string {
	return v.env
}

func NewViperConfig(env Environment, defaultConfigSrc string, configTypes ...ConfigSourceType) (AppConfig, error) {

	if env == EnvDefault {
		env = EnvDev
	}
	if len(configTypes) == 0 {
		configTypes = append(configTypes, ConfigSourceYaml)
	}
	for i := range configTypes {
		viper.SetConfigType(configTypes[i])
	}
	err := viper.ReadConfig(bytes.NewBuffer([]byte(defaultConfigSrc)))
	if err != nil {
		return nil, fmt.Errorf("failed to read config: %v", err)
	}
	v := viper.GetViper()
	return &ViperConfig{
		viper: v,
		env:   env,
	}, nil
}
