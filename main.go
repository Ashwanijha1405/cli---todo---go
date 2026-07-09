package main

import (
	"cli-todo/internal/cli"
	"cli-todo/internal/logger"
	"cli-todo/internal/repository/sqlite"
	"cli-todo/internal/todo"
	"fmt"
	"log"
	"os"
)

func main() {
	logger.Log.Info("application started")

	// Instantiate the repository
	repo, err := sqlite.NewSQLiteRepository("todos.db")
	if err != nil {
		log.Fatal(err)
	}

	// Inject into service
	todoService := todo.NewTodoService(repo)

	if len(os.Args) < 2 {
		fmt.Println("Please provide a command")
		return
	}

	command := os.Args[1]
	cli.HandleCommand(command, os.Args[2:], todoService)
}
