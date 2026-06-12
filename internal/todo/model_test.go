package todo

import "testing"

func TestTodoCreation(t *testing.T) {
	todo := Todo{
		ID:        1,
		Task:      "Learn Go",
		Completed: false,
	}

	if todo.ID != 1 {
		t.Errorf("expected ID 1, got %d", todo.ID)
	}

	if todo.Task != "Learn Go" {
		t.Errorf("unexpected task")
	}

	if todo.Completed {
		t.Errorf("todo should not be completed")
	}
}
