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

// getdeviceCmd represents the getdevice command
var getdeviceCmd = &cobra.Command{
	Use:   "getdevice",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println()
		if len(args) < 1 {
			fmt.Println("Please provide the AgentID of the device.")
			os.Exit(1)
		}
		device, err := apiClient.GetDevice(args[0])
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
	rootCmd.AddCommand(getdeviceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getdeviceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getdeviceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
