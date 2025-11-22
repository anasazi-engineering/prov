package cmd

import (
    "context"
    "fmt"
    "time"

    "github.com/spf13/cobra"
)

var userGetCmd = &cobra.Command{
    Use:   "get [id]",
    Short: "Get a user by ID",
    Args:  cobra.ExactArgs(1),
    RunE: func(cmd *cobra.Command, args []string) error {
        ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
        defer cancel()

        user, err := apiClient.GetUser(ctx, args[0])
        if err != nil {
            return err
        }

        fmt.Printf("ID: %s\nName: %s\nEmail: %s\n", user.ID, user.Name, user.Email)
        return nil
    },
}

func init() {
    userCmd.AddCommand(userGetCmd)
}
