package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/simenghe/knowntasker/taskswrapper"
	"github.com/spf13/cobra"
)

// listtCmd represents the listt command
var listtCmd = &cobra.Command{
	Use:   "listt",
	Short: "List the tasks of a single list. Use -number to specify the which list, default = 1",
	Run: func(cmd *cobra.Command, args []string) {
		srv, err := taskswrapper.GetTasksService()
		if err != nil {
			log.Fatal("Service Problem")
		}
		if len(args) > 0 {
			listIndex, err := strconv.Atoi(args[0])
			if err != nil {
				log.Fatalf("Input was not a valid number")
			}
			tasksList, err := taskswrapper.GetTaskList(srv, listIndex)
			if err != nil {
				log.Fatalln(err.Error())
			}
			items, err := taskswrapper.GetAllTaskListItems(srv, tasksList.Id)
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Printf("Displaying list : %s", tasksList.Title)
			for i, item := range items.Items {
				fmt.Printf("[%d] %s\n", i+1, item.Title)
			}
		} else {
			firstTaskList, err := taskswrapper.GetFirstTaskList(srv)
			if err != nil {
				log.Fatalf(err.Error())
			}
			fmt.Printf("Displaying list : %s", firstTaskList.Title)
			// Print the default list!
			items, err := taskswrapper.GetAllTaskListItems(srv, firstTaskList.Id)
			if err != nil {
				log.Fatalf(err.Error())
			}
			for i, item := range items.Items {
				fmt.Printf("[%d] %s\n", i+1, item.Title)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listtCmd)
}
