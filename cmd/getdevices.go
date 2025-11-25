/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
