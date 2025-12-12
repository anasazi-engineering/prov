/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// getbootboxCmd represents the getbootbox command
var getbootboxCmd = &cobra.Command{
	Use:   "getbootbox <agent ID>",
	Short: "Get specific device details",
	Long: `
Get details of a specific BootBox device by providing its 
AgentID. This command retrieves information such as the
device's friendly name, date joined, and last seen timestamp.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println()
		if len(args) < 1 {
			fmt.Println("Please provide the AgentID for the device.")
			os.Exit(1)
		}
		device, err := apiClient.GetBootbox(args[0])
		if err != nil {
			fmt.Printf("Error getting device: %v\n", err)
			return
		}
		fmt.Println("----------------------------------------------------------")
		fmt.Printf("AgentID:          %s\n", device.AgentID)
		fmt.Printf("Friendly Name:    %s\n", device.FriendlyName)
		fmt.Printf("Date Joined:      %s\n", time.Unix(device.CreatedAt, 0).Format(time.RFC822))
		fmt.Printf("Last Seen:        %s\n", time.Unix(device.LastSeen, 0).Format(time.RFC822))
		fmt.Println("----------------------------------------------------------")
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(getbootboxCmd)
}
