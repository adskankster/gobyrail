// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
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

	"github.com/adskankster/gobyrail/network"
	"github.com/adskankster/gobyrail/replies"
	"github.com/adskankster/gobyrail/requests"
	"github.com/spf13/cobra"
)

var cvs string
var cvs2 string
var timespan string

// stationCmd represents the station command
var stationCmd = &cobra.Command{
	Use:   "station",
	Short: "Lists the services to/from a station for a specified time period",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: nil,
}

// all represents the station all command
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "List all services to/from the station",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("station arr called")
		return nil
	},
}

// arr represents the station arr command
var arrCmd = &cobra.Command{
	Use:   "arr",
	Short: "List only arrivals at the station",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

// dep represents the station dep command
var depCmd = &cobra.Command{
	Use:   "dep",
	Short: "List only departures from the station",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		return getDepServices()
	},
}

// name represents the station name command
var nameCmd = &cobra.Command{
	Use:   "name",
	Short: "Look up the 3 letter CVS abbreviation for a station",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("station name called")
		return nil
	},
}

func init() {
	RootCmd.AddCommand(stationCmd)
	stationCmd.AddCommand(allCmd)
	stationCmd.AddCommand(arrCmd)
	stationCmd.AddCommand(depCmd)
	stationCmd.PersistentFlags().StringVarP(&cvs, "cvs", "c", "", "Station to look up")
	stationCmd.PersistentFlags().StringVarP(&timespan, "span", "s", "120", "Timespan in minutes (up to 120)")
	allCmd.LocalFlags().StringVarP(&cvs2, "tofrom", "t", "", "Show only services to/from this station.")
	arrCmd.LocalFlags().StringVarP(&cvs2, "from", "f", "", "Show only services from this station")
	depCmd.LocalFlags().StringVarP(&cvs2, "to", "t", "", "Show only services to this station")
}

//
func getDepServices() error {
	dbreq, err := requests.BuildGetDepartureBoardRequest(
		10, "SHF", "", "from", 0, 120,
	)
	if err != nil {
		return fmt.Errorf("Error getting request: %v", err)
	}

	b, err := network.MakeRequest("GetDepartureBoard", dbreq)
	if err != nil {
		return fmt.Errorf("Error building xml: %v", err)
	}

	// This is request dependant
	rep := replies.GetStationBoardResult{Template: "StationBoard"}

	// This is generic
	err = replies.HandleReply(b, &rep)
	if err != nil {
		return fmt.Errorf("Error handling reply xml: %v", err)
	}

	return nil
}
