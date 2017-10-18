// Copyright © 2016 Adam Wellings
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
	"io/ioutil"
	"os"

	"github.com/adskankster/gobyrail/helpers"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile       string
	debug         bool
	authtoken     string
	authtokenfile string
	format        string
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "gobyrail",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	RunE: func(cmd *cobra.Command, args []string) error {

	//		return nil
	//	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Printf("goByRail: %v\r\r", err)
		os.Exit(-1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/gobyrail.yaml)")
	RootCmd.PersistentFlags().StringVarP(&format, "format", "f", "", "Output format: fmt (default), xml, json")
	RootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", true, "Print debug messages to stdout")
	RootCmd.PersistentFlags().StringVarP(&authtoken, "token", "t", "", "Authorisation token")
	RootCmd.PersistentFlags().StringVarP(&authtokenfile, "tokenfile", "a", "", "File continaing authorisation token")

	cobra.OnInitialize(initConfig)
	cobra.OnInitialize(initOptions)
}

func initOptions() {
	helpers.SetDebug(debug)

	if format != "" {
		viper.Set("format", format)
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetConfigName("gobyrail")       // name of config file (without extension)
	viper.AddConfigPath("$HOME/.config/") // adding home directory as first search path

	viper.SetDefault("authtokenfile", "authtoken")
	viper.SetDefault("format", "fmt")

	viper.SetEnvPrefix("GBR_")
	viper.BindEnv("authtoken")
	viper.BindEnv("authtokenfile")
	viper.BindEnv("cfgfile")
	viper.AutomaticEnv()

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	// authtoken will override authtoken file

	// TODO - improve error handling here
	if authtokenfile != "" {
		authtoken, err := ioutil.ReadFile(authtokenfile)
		if err == nil {
			viper.Set("authtoken", authtoken)
		} else {
			fmt.Println("Warning: error reading token file.")
		}
	}

	if authtoken != "" {
		viper.Set("authtoken", authtoken)
	}
}
