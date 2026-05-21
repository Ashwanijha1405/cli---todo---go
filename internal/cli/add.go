package cli

import (
	"cli-todo/internal/storage"
	"cli-todo/internal/todo"
	"fmt"
)

func HandleAdd(args []string, todos []todo.Todo) []todo.Todo {

	// Validation Layer
	if len(args) < 1 {
		fmt.Println("Please provide todo title")
		return todos
	}

	// Command Input
	title := args[0]

	// Business Logic Layer
	todos = todo.AddTodo(todos, title)

	// Persistence Layer
	storage.SaveTodos(todos)

	// Presentation Layer
	todo.ListTodos(todos)

	return todos
}
