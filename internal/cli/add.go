package cli

import (
	"cli-todo/internal/constants"
	"cli-todo/internal/logger"
	"cli-todo/internal/todo"
	"fmt"
)

func HandleAdd(args []string, todoService *todo.TodoService) {

	// Validation Layer
	if len(args) < 1 {

		logger.Log.Error("missing todo title for add command")

		fmt.Println(constants.MsgProvideTodoID)
		return
	}

	// Command Input
	title := args[0]

	logger.Log.Info(
		"processing add command",
		"title", title,
	)

	// Business Logic Layer
	_, err := todoService.AddTodo(title)
	if err != nil {
		fmt.Println("Error saving todos:", err)
		return
	}

	logger.Log.Info(
		"add command executed successfully",
		"title", title,
	)

	// Presentation Layer
	_, err = todoService.ListTodos()
	if err != nil {
		fmt.Println("Error listing todos:", err)
	}
}