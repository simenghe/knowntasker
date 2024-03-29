/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/simenghe/knowntasker/taskswrapper"
	"github.com/spf13/cobra"
)

// listlCmd represents the listl command
var listlCmd = &cobra.Command{
	Use:   "listl",
	Short: "Shows all available Tasks Lists of the current user",
	Run: func(cmd *cobra.Command, args []string) {
		srv, err := taskswrapper.GetTasksService()
		if err != nil {
			log.Fatal("Service Problem")
		}
		userLists, err := taskswrapper.GetAllTasksLists(srv)
		for i, list := range userLists.Items {
			fmt.Printf("[%d] %s\n", i+1, list.Title)
		}
		if err != nil {
			log.Fatalf("Problem fetching UserLists")
		}
	},
}

func init() {
	rootCmd.AddCommand(listlCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listlCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
