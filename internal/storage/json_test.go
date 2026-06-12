package storage

import (
	"path/filepath"
	"testing"

	"github.com/roystonmoraes/go-todo-cli/internal/todo"
)

func TestSaveAndLoadTodos(t *testing.T) {
	tempDir := t.TempDir()
	filepath := filepath.Join(tempDir, "todos.json")
	repo := NewJSONRepository(filepath)
	expected := []todo.Todo{
		{
			ID:        1,
			Task:      "Learn Go",
			Completed: false,
		},
		{
			ID:        2,
			Task:      "Build CLI",
			Completed: true,
		},
	}

	if err := repo.Save(expected); err != nil {
		t.Fatalf("failed to save todos: %v", err)
	}

	actual, err := repo.Load()
	if err != nil {
		t.Fatalf("failed to load todos: %v", err)
	}

	if len(actual) != len(expected) {
		t.Fatalf("expected %d todos, got %d", len(expected), len(actual))
	}

	for i := range expected {
		if expected[i] != actual[i] {
			t.Errorf("expected %+v, got %+v", expected[i], actual[i])
		}
	}
}

func TestLoadMissingFileReturnsEmptySlice(t *testing.T) {
	tempDir := t.TempDir()

	filePath := filepath.Join(tempDir, "missing.json")

	repo := NewJSONRepository(filePath)

	todos, err := repo.Load()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(todos) != 0 {
		t.Fatalf("expected empty slice")
	}
}

func TestLoadEmptyFile(t *testing.T) {
	tempDir := t.TempDir()

	filePath := filepath.Join(tempDir, "empty.json")

	repo := NewJSONRepository(filePath)

	if err := repo.Save([]todo.Todo{}); err != nil {
		t.Fatalf("save failed: %v", err)
	}

	todos, err := repo.Load()
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}

	if len(todos) != 0 {
		t.Fatalf("expected empty slice")
	}
}
