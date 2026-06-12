package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all todos",
	Run: func(cmd *cobra.Command, args []string) {
		// use global service

		todos, err := service.List()
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		if len(todos) == 0 {
			fmt.Println("No todos found.")
			return
		}

		for i, t := range todos {
			status := " "
			if t.Completed {
				status = "x"
			}

			fmt.Printf("%d. [%s] %s\n", i+1, status, t.Task)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
