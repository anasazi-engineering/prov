/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"log"
	"time"

	"github.com/spf13/cobra"
)

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "logout user",
	Long: `
Logout is used to logout out the current user. Access
and Refresh tokens will be removed from the local storage.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		// Call login API
		err := apiClient.Logout(ctx)
		if err != nil {
			log.Fatalf("User logout failed: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}
