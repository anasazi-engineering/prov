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
	"time"

	"github.com/spf13/cobra"
)

// getdevicesCmd represents the getdevices command
var getdevicesCmd = &cobra.Command{
	Use:   "getdevices",
	Short: "List authorized devices",
	Long: `List devices that have been authorized to
connect the server. Devices may or may not
have been assigned a recipe yet.`,
	Run: func(cmd *cobra.Command, args []string) {
		devices, err := apiClient.GetDevices()
		if err != nil {
			fmt.Printf("Error getting devices: %v\n", err)
			return
		}
		fmt.Println("\n  AgentID                             Recipe          %    Date Joined")
		fmt.Println("-------------------------------------------------------------------------------")
		for _, device := range devices {
			t := time.Unix(device.CreatedAt, 0).Format(time.RFC822)
			fmt.Printf("%s%16.14s%4.2d%22.19s\n",
				device.AgentID, device.AssdRecipe, device.RecipeProgress, t)
		}
		fmt.Println("-------------------------------------------------------------------------------")
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(getdevicesCmd)
}
