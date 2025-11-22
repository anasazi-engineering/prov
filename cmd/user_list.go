package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var userListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all users",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Taking over user list command...")

		fmt.Printf("Config file: %s\n", viper.ConfigFileUsed())

		// Print current config values
		fmt.Printf("Base URL: %s\n", viper.GetString("base_url"))
		fmt.Printf("Access Token: %s\n", viper.GetString("access_token"))
		fmt.Printf("Refresh Token: %s\n", viper.GetString("refresh_token"))

		//config.cfg.AccessToken = "Hello World"
		viper.Set("access_token", "Hello World")
		viper.WriteConfig()

		fmt.Printf("Base URL: %s\n", viper.GetString("base_url"))
		fmt.Printf("Access Token: %s\n", viper.GetString("access_token"))
		fmt.Printf("Refresh Token: %s\n", viper.GetString("refresh_token"))

		return nil
	},
}

func init() {
	userCmd.AddCommand(userListCmd)
}
