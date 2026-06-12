package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a todo",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("invalid id")
			return
		}

		// use global service

		err = service.Delete(id)
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		fmt.Println("todo deleted:", id)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
