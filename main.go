package main

import (
	"cli-todo/internal/cli"
	"cli-todo/internal/constants"
	"cli-todo/internal/logger"
	"cli-todo/internal/storage"
	"cli-todo/internal/todo"
	"fmt"
	"os"
)

func main() {

	logger.Log.Info("application started")

	// Instantiate the repository (low-level detail)
	repo := storage.NewJSONRepository(constants.TodosFile)

	// Inject the repository into the service (high-level logic)
	todoService := todo.NewTodoService(repo)

	if len(os.Args) < 2 {
		fmt.Println("Please provide a command")
		return
	}

	command := os.Args[1]

	cli.HandleCommand(command, os.Args[2:], todoService)
}
