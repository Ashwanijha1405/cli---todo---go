package test

import (
	"testing"

	"cli-todo/internal/repository/contract"
)

func RunRepositoryTests(
	t *testing.T,
	repo contract.TodoRepository,
) {

	t.Run("Create Todo", func(t *testing.T) {
	todo, err := repo.Create("Learn Testing")

	if err != nil {
		t.Fatalf("Create() returned an error: %v", err)
	}

	if todo.Title != "Learn Testing" {
		t.Fatalf("expected title %q, got %q", "Learn Testing", todo.Title)
	}

	if todo.Completed {
		t.Fatalf("new todo should not be completed")
	}

	if todo.ID == 0 {
		t.Fatalf("expected non-zero ID")
	}

  })

  t.Run("GetAll Todos", func(t *testing.T) {

	_, err := repo.Create("Todo 1")
	if err != nil {
		t.Fatalf("Create() returned an error: %v", err)
	}

	_, err = repo.Create("Todo 2")
	if err != nil {
		t.Fatalf("Create() returned an error: %v", err)
	}

	todos, err := repo.GetAll()
	if err != nil {
		t.Fatalf("GetAll() returned an error: %v", err)
	}

	if len(todos) < 2 {
		t.Fatalf("expected at least 2 todos, got %d", len(todos))
	}

	foundTodo1 := false
	foundTodo2 := false

	for _, todo := range todos {
		switch todo.Title {
		case "Todo 1":
			foundTodo1 = true
		case "Todo 2":
			foundTodo2 = true
		}
	}


	if !foundTodo1 {
		t.Fatal("Todo 1 not found")
	}

	if !foundTodo2 {
		t.Fatal("Todo 2 not found")
	}

  })

  t.Run("Update Todo", func(t *testing.T) {
	todo, err := repo.Create("Old Title")
	if err != nil {
		t.Fatalf("Create() returned an error: %v", err)
	}

	todo.Title = "New Title"
	todo.Completed = true

	err = repo.Update(todo)
	if err != nil {
		t.Fatalf("Update() returned an error: %v", err)
	}

	todos, err := repo.GetAll()
	if err != nil {
		t.Fatalf("GetAll() returned an error: %v", err)
	}

	var updated *contract.Todo

	for i := range todos {
		if todos[i].ID == todo.ID {
			updated = &todos[i]
			break
		}
	}

	if updated == nil {
		t.Fatal("updated todo not found")
	}

	if updated.Title != "New Title" {
		t.Fatalf("expected title %q, got %q", "New Title", updated.Title)
	}

	if !updated.Completed {
		t.Fatal("expected todo to be completed")
	}

  })

  t.Run("Delete Todo", func(t *testing.T) {
	todo, err := repo.Create("Delete Me")
	if err != nil {
		t.Fatalf("Create() returned an error: %v", err)
	}

	err = repo.Delete(todo.ID)
	if err != nil {
		t.Fatalf("Delete() returned an error: %v", err)
	}

	todos, err := repo.GetAll()
	if err != nil {
		t.Fatalf("GetAll() returned an error: %v", err)
	}

	for _, item := range todos {
		if item.ID == todo.ID {
			t.Fatal("todo was not deleted")
		}
	}

  })

  t.Run("Update Missing Todo", func(t *testing.T) {
	fakeTodo := contract.Todo{
		ID:        9999,
		Title:     "Doesn't Exist",
		Completed: false,
	}

	err := repo.Update(fakeTodo)
	if err == nil {
		t.Fatal("expected Update() to return an error for a missing todo")
	}

  })

  t.Run("Delete Missing Todo", func(t *testing.T) {
	err := repo.Delete(9999)

	if err == nil {
		t.Fatal("expected Delete() to return an error for a missing todo")
	}

  })

}