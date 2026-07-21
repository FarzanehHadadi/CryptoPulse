package config

import (
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func Load() (*Config, error) {

	_ = godotenv.Load()

	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./configs")

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()
	v.SetConfigName("config")
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	v.SetConfigName("scheduler")
	if err := v.MergeInConfig(); err != nil {
		return nil, err
	}

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config

	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
