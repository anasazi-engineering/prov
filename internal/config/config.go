package config

import (
    "errors"
    "github.com/spf13/viper"
)

type Config struct {
    BaseURL string `mapstructure:"base_url"`
    Token   string `mapstructure:"token"`
}

func Load(cfgFile string) (*Config, error) {
    var cfg Config
    if err := viper.Unmarshal(&cfg); err != nil {
        return nil, err
    }
    if cfg.BaseURL == "" {
        return nil, errors.New("base_url must be configured")
    }
    return &cfg, nil
}
