package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [task]",
	Short: "Add a new todo",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		task := args[0]

		// use global service

		err := service.Add(task)
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		fmt.Println("todo added:", task)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
