package cmd

import (
    "context"
    "fmt"
    "time"

    "github.com/spf13/cobra"
)

var userListCmd = &cobra.Command{
    Use:   "list",
    Short: "List all users",
    RunE: func(cmd *cobra.Command, args []string) error {
        ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
        defer cancel()

        users, err := apiClient.ListUsers(ctx)
        if err != nil {
            return err
        }

        fmt.Println("Users:")
        for _, u := range users {
            fmt.Printf(" - %s (%s)\n", u.Name, u.ID)
        }
        return nil
    },
}

func init() {
    userCmd.AddCommand(userListCmd)
}
