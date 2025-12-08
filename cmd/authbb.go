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
