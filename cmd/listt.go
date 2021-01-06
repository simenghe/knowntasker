package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listtCmd represents the listt command
var listtCmd = &cobra.Command{
	Use:   "listt",
	Short: "List the tasks of a single list. Use -number to specify the which list, default = 1",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("listt called")
	},
}

func init() {
	rootCmd.AddCommand(listtCmd)
}
