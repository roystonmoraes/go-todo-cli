package cmd

import (
	"fmt"
	"strconv"

	"github.com/roystonmoraes/go-todo-cli/internal/storage"
	"github.com/roystonmoraes/go-todo-cli/internal/todo"
	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete [id]",
	Short: "Mark a todo as completed",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("invalid id")
			return
		}

		repo := storage.NewJSONRepository("data/todos.json")
		service := todo.NewService(repo)

		err = service.Complete(id)
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		fmt.Println("todo completed:", id)
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
