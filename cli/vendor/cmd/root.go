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
	"entity"
	"os"
	"utils"

	errors "github.com/go-errors/errors"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	logger  = utils.NewLogger()
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "agenda",
	Short: "welcome to Agenda",
	Long:  `welcome to Agenda`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	defer utils.CloseLogFile()
	if _, present := os.LookupEnv("DEBUG"); present == false {
		defer func() {
			if err := recover(); err != nil {
				logger.Fatalln("[execute]", err)
			}
			return
		}()
	}
	if err := RootCmd.Execute(); err != nil {
		return
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			logger.Fatalln("Failed to find directory", errors.Wrap(err, 0))
		}

		// Search config in home directory with name ".agenda-go" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".agenda-go")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	var err error
	err = viper.ReadInConfig()

	// set current working directory
	c := viper.GetString("c")
	if len(c) > 0 {
		err = os.Chdir(c)
		if err != nil {
			logger.Fatalf("Failed to change current working directory: %v", err)
		}
	}
	// set log file
	logPath := viper.GetString("log")
	utils.InitLogFile(logPath)

	if err == nil {
		logger.Println("[init] using configs from config file:", viper.ConfigFileUsed())
	} else {
		logger.Println("[init] using default configs")
	}
	entity.Init()
}
