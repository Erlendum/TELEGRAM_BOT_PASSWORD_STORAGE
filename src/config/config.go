package config

import (
	"github.com/spf13/viper"
	"src/cmd/modes/flags"
)

type Config struct {
	Redis flags.RedisFlags       `mapstructure:"redis"`
	Bot   flags.TelegramBotFlags `mapstructure:"bot"`
}

func (c *Config) ParseConfig(configFileName, pathToConfig string) error {
	v := viper.New()
	v.SetConfigName(configFileName)
	v.SetConfigType("json")
	v.AddConfigPath(pathToConfig)

	err := v.ReadInConfig()
	if err != nil {
		return err
	}

	err = v.Unmarshal(c)
	if err != nil {
		return err
	}

	return nil
}
