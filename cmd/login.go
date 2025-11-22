/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"syscall"

	"golang.org/x/term"

	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to Provisioner API server",
	Long: `
The 'login' subcommand is used to login to Provisioner API server.
The user will be prompted to enter their username, pasword, and 
2FA token from their Authenticator app. On success, a session token
will be stored in the user's config file for future requests.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("login called")

		var username, password string
		// Get Username
		fmt.Print("Enter Username: ")
		_, err := fmt.Scanln(&username)
		if err != nil {
			// Ignore EOF (Ctrl+D) during username read, treat it as a normal error
			if err.Error() == "unexpected newline" || err.Error() == "EOF" {
				log.Fatalf("failed to read username: %w", err)
			}
			// Clear the input buffer if Scanln didn't read an entire line
			// (though in this case, we just exit on error)
			log.Fatalf("failed to read username: %w", err)
		}

		// Get Password without echoing input
		fmt.Print("Enter Password: ")

		// Get the file descriptor for standard input
		fd := int(syscall.Stdin)

		// Read the password using golang.org/x/term.ReadPassword.
		// This function turns off terminal echoing so the input is hidden.
		bytePassword, err := term.ReadPassword(fd)
		if err != nil {
			log.Fatalf("failed to read password: %w", err)
		}

		// Print a newline after reading the password since ReadPassword doesn't add one
		fmt.Println()

		password = string(bytePassword)
		fmt.Printf("Username: %s\n", username)
		fmt.Printf("Password: %s\n", password)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
