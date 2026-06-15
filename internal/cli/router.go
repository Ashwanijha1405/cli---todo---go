package cli

import (
	"cli-todo/internal/constants"
	"cli-todo/internal/todo"
	"cli-todo/internal/constants"
	"fmt"
)

func HandleCommand(command string, args []string, todoService *todo.TodoService) {

	switch command {

	case constants.CommandAdd:
		HandleAdd(args, todoService)

	case constants.CommandList:
		HandleList(todoService)

	case constants.CommandDone:
		HandleDone(args, todoService)

	case constants.CommandDelete:
		HandleDelete(args, todoService)

	case constants.CommandHelp:
		PrintHelp()

	default:
		fmt.Println("Unknown command:", command)
		PrintHelp()
	}
}
