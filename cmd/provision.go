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
	"log"
	"time"

	"github.com/spf13/cobra"
)

// provisionCmd represents the provision command
var provisionCmd = &cobra.Command{
	Use:   "provision",
	Short: "Apply recipe to node",
	Long: `The provision command is used to select a recipe
that will be used to provision the Worker node specified.
The command requires the Worker's agentID and the URL
for the recipe to be applied.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		// Get command line flags
		agentID, _ := cmd.Flags().GetString("worker")
		provURL, _ := cmd.Flags().GetString("url")
		fmt.Printf("Agent ID: %s\n", agentID)
		fmt.Printf("Recipe URL: %s\n", provURL)

		err := apiClient.ApplyRecipe(ctx, agentID, provURL)
		if err != nil {
			log.Fatalf("Error provisioning worker: %v\n", err)
			return
		}
		fmt.Println("Agent provisioned successfully.")
	},
}

func init() {
	rootCmd.AddCommand(provisionCmd)
	provisionCmd.Flags().StringP("worker", "w", "", "Agent ID for Worker to provision")
	provisionCmd.Flags().StringP("url", "u", "", "URL for recipe")

	provisionCmd.MarkFlagRequired("worker")
	provisionCmd.MarkFlagRequired("url")
}
