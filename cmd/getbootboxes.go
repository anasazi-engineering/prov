/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// getbootboxesCmd represents the getbootboxes command
var getbootboxesCmd = &cobra.Command{
	Use:   "getbootboxes",
	Short: "List authorized BootBox devices",
	Long: `List BootBoxes that have been authorized to
connect the server.`,
	Run: func(cmd *cobra.Command, args []string) {
		devices, err := apiClient.GetBootboxes()
		if err != nil {
			fmt.Printf("Error getting devices: %v\n", err)
			return
		}
		fmt.Println("\n             AgentID                           Date Joined")
		fmt.Println("--------------------------------------------------------------")
		for _, device := range devices {
			t := time.Unix(device.CreatedAt, 0).Format(time.RFC822)
			fmt.Printf("%s%26.19s\n", device.AgentID, t)
		}
		fmt.Println("--------------------------------------------------------------")
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(getbootboxesCmd)
}
