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
	Use:   "prov",
	Short: "CLI for ProvisionerAPI server",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load(cfgFile)
		if err != nil {
			return err
		}
		apiClient = api.NewClient(cfg.BaseURL, cfg.Token)
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
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file ($HOME/.provcli)")
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
