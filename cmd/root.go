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
package cmd

import (
	"fmt"
	"os"

	"prov/internal/api"
	"prov/internal/config"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile   string
	apiClient api.Client
)

var rootCmd = &cobra.Command{
	Use:   "prov <command>",
	Short: "CLI for ProvisionerAPI server",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load(cfgFile)
		if err != nil {
			return err
		}
		apiClient = api.NewClient(cfg.BaseURL, config.Tokens{
			AccessToken:  cfg.AccessToken,
			RefreshToken: cfg.RefreshToken,
		})
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, _ := os.UserHomeDir()
		viper.AddConfigPath(home)
		viper.SetConfigName(".provcli")
		viper.SetConfigType("json")
	}

	viper.SetEnvPrefix("PROVCLI")
	viper.AutomaticEnv()

	viper.ReadInConfig()
}
