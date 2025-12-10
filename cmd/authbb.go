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

		// Validate that an OTP argument is provided
		if len(args) < 1 {
			fmt.Println("\nError: OTP argument is required\n")
			// output usage information
			cmd.Usage()
			return
		} else if len(args) > 1 {
			fmt.Println("\nError: Too many arguments provided. Only the OTP is required.\n")
			// output usage information
			cmd.Usage()
			return
		}

		// Call the API client to authorize the BootBox with the provided OTP
		err := apiClient.AuthBootBox(ctx, args[0])
		if err != nil {
			fmt.Printf("Error authorizing BootBox: %v\n", err)
			return
		}
		fmt.Printf("OTP: %s authorized\n\n", args[0])

	},
}

func init() {
	rootCmd.AddCommand(authbbCmd)
}
