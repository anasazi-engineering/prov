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
	"time"

	"github.com/spf13/cobra"
)

// getworkerCmd represents the getworker command
var getworkerCmd = &cobra.Command{
	Use:   "getworker <agent ID>",
	Short: "Get specific device details",
	Long: `Get details of a specific Worker device by providing its 
AgentID. This command retrieves information such as the
device's friendly name, assigned recipe, recipe progress,
date joined, and last seen timestamp.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println()
		if len(args) < 1 {
			fmt.Println("Please provide the AgentID for the device.")
			os.Exit(1)
		}
		device, err := apiClient.GetWorker(args[0])
		if err != nil {
			fmt.Printf("Error getting device: %v\n", err)
			return
		}
		fmt.Println("----------------------------------------------------------")
		fmt.Printf("AgentID:          %s\n", device.AgentID)
		fmt.Printf("Friendly Name:    %s\n", device.FriendlyName)
		fmt.Printf("Assigned Recipe:  %s\n", device.AssdRecipe)
		fmt.Printf("Recipe Progress:  %d%%\n", device.RecipeProgress)
		fmt.Printf("Date Joined:      %s\n", time.Unix(device.CreatedAt, 0).Format(time.RFC822))
		fmt.Printf("Last Seen:        %s\n", time.Unix(device.LastSeen, 0).Format(time.RFC822))
		fmt.Println("----------------------------------------------------------")
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(getworkerCmd)
}
