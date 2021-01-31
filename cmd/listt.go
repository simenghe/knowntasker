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
		fmt.Println("listt called")
		if len(args) > 0 {
			listIndex, err := strconv.Atoi(args[0])
			if err != nil {
				log.Fatalf("Input was not a valid number")
			}
			fmt.Println(listIndex)
			tasksList, err := taskswrapper.GetTaskList(srv, listIndex)
			if err != nil {
				log.Fatalln(err)
			}
			items, err := taskswrapper.GetAllTaskListItems(srv, tasksList.Id)
			if err != nil {
				log.Fatalln(err)
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
