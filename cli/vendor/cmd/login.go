package cmd

import (
	"fmt"
	"os"
	"service"

	"github.com/spf13/cobra"
)

var logincommand = &cobra.Command{
	Use:   "login",
	Short: "log in Agenda",
	Long:  `Log in Agenda with username and password`,
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		var err error
		err = service.Login(username, password)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
		} else {
			fmt.Println("Welcome, ", username)
		}
	},
}

func init() {
	RootCmd.AddCommand(logincommand)
	logincommand.Flags().StringP("username", "u", "", "username")
	logincommand.Flags().StringP("password", "p", "", "password")
}
