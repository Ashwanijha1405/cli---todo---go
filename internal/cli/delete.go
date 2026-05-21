package cli

import (
	"cli-todo/internal/storage"
	"cli-todo/internal/todo"
	"fmt"
	"strconv"
)

func HandleDelete(args []string, todos []todo.Todo) []todo.Todo {

	// Validation Layer
	if len(args) < 1 {
		fmt.Println("Please provide todo ID")
		return todos
	}

	// Input Parsing
	id, err := strconv.Atoi(args[0])

	if err != nil {
		fmt.Println("Invalid todo ID")
		return todos
	}

	// Business Logic Layer
	todos = todo.DeleteTodo(todos, id)

	// Persistence Layer
	storage.SaveTodos(todos)

	return todos
}
