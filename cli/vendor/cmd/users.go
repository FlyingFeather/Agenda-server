package cmd

import (
	"fmt"
	"os"
	"service"

	"github.com/spf13/cobra"
)

var listCommand = &cobra.Command{
	Use:   "list users",
	Short: "list all registered users",
	Long:  `list all registered users`,
	Run: func(cmd *cobra.Command, args []string) {
		results, err := service.GetAllUsers()
		if err == nil {
			fmt.Println("List all registered users")
			for i, user := range results {
				fmt.Printf("%d. \n", i)
				fmt.Printf("Username: %s\n", user.Username)
				fmt.Printf("Email: %s\n", user.Email)
				fmt.Printf("Telephone: %s\n\n", user.Telephone)
			}
		} else {
			fmt.Fprintln(os.Stderr, "Error:", err)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCommand)
}
