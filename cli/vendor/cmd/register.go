// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"
	"service"

	"github.com/spf13/cobra"
)

var registercommand = &cobra.Command{
	Use:   "register",
	Short: "now please register an account",
	Long:  `Now please register an account with username, password, email and telephone`,
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		email, _ := cmd.Flags().GetString("email")
		telephone, _ := cmd.Flags().GetString("telephone")
		var err error
		err = service.Register(username, password, email, telephone)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
		} else {
			fmt.Println("Registered user:", username)
		}
	},
}

func init() {
	RootCmd.AddCommand(registercommand)
	registercommand.Flags().StringP("username", "u", "", "username")
	registercommand.Flags().StringP("password", "p", "", "password")
	registercommand.Flags().StringP("email", "m", "", "your email")
	registercommand.Flags().StringP("telephone", "t", "", "your telephone number")
}
