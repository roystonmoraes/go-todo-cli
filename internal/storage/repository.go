package storage

import "github.com/roystonmoraes/go-todo-cli/internal/todo"

type Repository interface {
	Load() ([]todo.Todo, error)
	Save([]todo.Todo) error
}
