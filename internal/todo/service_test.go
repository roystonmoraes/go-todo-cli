package todo

import "testing"

type mockRepository struct {
	todos []Todo
}

func (m *mockRepository) Load() ([]Todo, error) {
	return m.todos, nil
}

func (m *mockRepository) Save(todos []Todo) error {
	m.todos = todos
	return nil
}

func TestAddTodo(t *testing.T) {
	repo := &mockRepository{}
	service := NewService(repo)

	err := service.Add("Learn Go")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(repo.todos) != 1 {
		t.Fatalf("expected 1 todo, got %d", len(repo.todos))
	}

	if repo.todos[0].Task != "Learn Go" {
		t.Fatalf("unexpected task")
	}

	if repo.todos[0].ID != 1 {
		t.Fatalf("expected id 1")
	}
}

func TestListTodos(t *testing.T) {
	repo := &mockRepository{
		todos: []Todo{
			{
				ID:   1,
				Task: "Learn Go",
			},
		},
	}

	service := NewService(repo)

	todos, err := service.List()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(todos) != 1 {
		t.Fatalf("expected 1 todo")
	}
}

func TestCompleteTodo(t *testing.T) {
	repo := &mockRepository{
		todos: []Todo{
			{
				ID:        1,
				Task:      "Learn Go",
				Completed: false,
			},
		},
	}

	service := NewService(repo)

	err := service.Complete(1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !repo.todos[0].Completed {
		t.Fatalf("todo should be completed")
	}
}

func TestCompleteMissingTodo(t *testing.T) {
	repo := &mockRepository{}

	service := NewService(repo)

	err := service.Complete(99)

	if err == nil {
		t.Fatal("expected error")
	}
}

func TestDeleteTodo(t *testing.T) {
	repo := &mockRepository{
		todos: []Todo{
			{
				ID:   1,
				Task: "Learn Go",
			},
		},
	}

	service := NewService(repo)

	err := service.Delete(1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(repo.todos) != 0 {
		t.Fatalf("expected todo to be deleted")
	}
}

func TestDeleteMissingTodo(t *testing.T) {
	repo := &mockRepository{}

	service := NewService(repo)

	err := service.Delete(123)

	if err == nil {
		t.Fatal("expected error")
	}
}
