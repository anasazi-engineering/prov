/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

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
		fmt.Println("logout called")
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}
