/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"prov/internal/api"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to Provisioner API server",
	Long: `
The 'login' subcommand is used to login to Provisioner API server.
The user will be prompted to enter their username, pasword, and 
2FA token from their Authenticator app. On success, a session token
will be stored in the user's config file for future requests.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()
		var creds api.Credentials
		creds.Username, _ = cmd.Flags().GetString("username")
		creds.Password, _ = cmd.Flags().GetString("password")
		creds.OrgID, _ = cmd.Flags().GetString("org-id")
		if creds.Username == "" || creds.Password == "" || creds.OrgID == "" {
			fmt.Println("\n** username, password, and org-id are required **\n")
			cmd.Help()
			os.Exit(0)
		}

		// Call login API
		tokens, err := apiClient.Login(ctx, creds)
		if err != nil {
			log.Fatalf("User login failed: %v\n", err)
		}

		// Store tokens in config
		viper.Set("access_token", tokens.AccessToken)
		viper.Set("refresh_token", tokens.RefreshToken)
		viper.WriteConfig()

		fmt.Println("User logged in successfully.")
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Add flags for username and password
	loginCmd.Flags().StringP("username", "u", "", "Username for login")
	loginCmd.Flags().StringP("password", "p", "", "Password for login")
	loginCmd.Flags().StringP("org-id", "o", "", "Organization ID for login")
}
