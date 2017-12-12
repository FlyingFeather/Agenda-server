package cmd

import (
	"fmt"
	"os"
	"service"

	"github.com/spf13/cobra"
)

var logoutcommand = &cobra.Command{
	Use:   "logout",
	Short: "Log out of Agenda",
	Long:  "Log out of Agenda",
	Run: func(cmd *cobra.Command, args []string) {
		err := service.Logout()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
		} else {
			fmt.Println("You have logged out successfully")
		}
	},
}

func init() {
	RootCmd.AddCommand(logoutcommand)
}
