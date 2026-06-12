package storage

import (
	"encoding/json"
	"os"

	"github.com/roystonmoraes/go-todo-cli/internal/todo"
)

type JSONRepository struct {
	FilePath string
}

func NewJSONRepository(path string) *JSONRepository {
	return &JSONRepository{
		FilePath: path,
	}
}

func (r *JSONRepository) Load() ([]todo.Todo, error) {
	if _, err := os.Stat(r.FilePath); os.IsNotExist(err) {
		return []todo.Todo{}, nil
	}

	data, err := os.ReadFile(r.FilePath)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return []todo.Todo{}, nil
	}

	var todos []todo.Todo

	if err := json.Unmarshal(data, &todos); err != nil {
		return nil, err
	}

	return todos, nil
}

func (r *JSONRepository) Save(todos []todo.Todo) error {
	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(r.FilePath, data, 0644)
}
