package cli

import (
	"cli-todo/internal/logger"
	"cli-todo/internal/constants"
	"cli-todo/internal/storage"
	"cli-todo/internal/todo"
	"fmt"
)

func HandleAdd(args []string, todos []todo.Todo) []todo.Todo {

	// Validation Layer
	if len(args) < 1 {

		logger.Log.Error("missing todo title for add command")

		fmt.Println(constants.MsgProvideTodoID)
		return todos
	}

	// Command Input
	title := args[0]

	logger.Log.Info(
		"processing add command",
		"title", title,
	)

	// Business Logic Layer
	todos = todo.AddTodo(todos, title)

	// Persistence Layer
	err := storage.SaveTodos(todos)

	if err != nil {

		logger.Log.Error(
			"failed to save todos after add",
			"error", err,
		)

		fmt.Println("Error saving todos:", err)
		return todos
	}

	logger.Log.Info(
		"add command executed successfully",
		"title", title,
	)

	// Presentation Layer
	todo.ListTodos(todos)

	return todos
}