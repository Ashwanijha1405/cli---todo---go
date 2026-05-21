package cli

import (
	"cli-todo/internal/todo"
	"fmt"
)

func HandleCommand(command string, args []string, todos []todo.Todo) []todo.Todo {

	switch command {

	case "add":
		return HandleAdd(args, todos)

	case "list":
		return HandleList(todos)

	case "done":
		return HandleDone(args, todos)

	case "delete":
		return HandleDelete(args, todos)

	case "help":
		PrintHelp()
		return todos

	default:
		fmt.Println("Unknown command:", command)
		PrintHelp()
		return todos
	}
}
