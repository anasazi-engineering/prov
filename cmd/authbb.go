/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// authbbCmd represents the authbb command
var authbbCmd = &cobra.Command{
	Use:   "authbb <OTP>",
	Short: "Authorizes a BootBox using the provided OTP",
	Long: `Authorizes a BootBox using the provided OTP. The OTP is
generally found on the BootBox device's display screen
or on the command line during debug.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
		defer cancel()
		fmt.Println("authbb called")
		err := apiClient.AuthBootBox(ctx, args[0])
		if err != nil {
			fmt.Printf("Error gauthorizing BootBox: %v\n", err)
			return
		}

	},
}

func init() {
	rootCmd.AddCommand(authbbCmd)
}
