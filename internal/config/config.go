/*
 * Anasazi Precision Engineering LLC CONFIDENTIAL
 *
 * Unpublished Copyright (c) 2025 Anasazi Precision Engineering LLC. All Rights Reserved.
 *
 * Proprietary to Anasazi Precision Engineering LLC and may be covered by patents, patents
 * in process, and trade secret or copyright law. Dissemination of this information or
 * reproduction of this material is strictly forbidden unless prior written
 * permission is obtained from Anasazi Precision Engineering LLC.
 */
package config

import (
	"errors"

	"github.com/spf13/viper"
)

type Config struct {
	BaseURL      string `mapstructure:"base_url"`
	AccessToken  string `mapstructure:"access_token"`
	RefreshToken string `mapstructure:"refresh_token"`
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
