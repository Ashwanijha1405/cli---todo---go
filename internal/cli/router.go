package cli

import (
	"cli-todo/internal/todo"
	"cli-todo/internal/constants"
	"fmt"
)

func HandleCommand(command string, args []string, todos []todo.Todo) []todo.Todo {

	switch command {

	case constants.CommandAdd:
		return HandleAdd(args, todos)

	case constants.CommandList:
		return HandleList(todos)

	case constants.CommandDone:
		return HandleDone(args, todos)

	case constants.CommandDelete:
		return HandleDelete(args, todos)

	case constants.CommandHelp:
		PrintHelp()
		return todos

	default:
		fmt.Println("Unknown command:", command)
		PrintHelp()
		return todos
	}
}
