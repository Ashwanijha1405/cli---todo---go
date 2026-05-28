package main

import (
	"cli-todo/internal/cli"
	"cli-todo/internal/storage"
	"fmt"
	"os"
)

func main() {

	todos, err := storage.LoadTodos()

    if err != nil {
	    fmt.Println("Error loading todos:", err)
	    return
    }

	if len(os.Args) < 2 {
		fmt.Println("Please provide a command")
		return
	}

	command := os.Args[1]

	todos = cli.HandleCommand(command, os.Args[2:], todos)

}
