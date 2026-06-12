package cmd

import (
	"github.com/roystonmoraes/go-todo-cli/internal/storage"
	"github.com/roystonmoraes/go-todo-cli/internal/todo"
	"github.com/spf13/cobra"
)

var service *todo.Service

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "Todo CLI application",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	repo := storage.NewJSONRepository("data/todos.json")
	service = todo.NewService(repo)
}
