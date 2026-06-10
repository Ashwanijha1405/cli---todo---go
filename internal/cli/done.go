package cli

import (
	"cli-todo/internal/storage"
	"cli-todo/internal/constants"
	"cli-todo/internal/todo"
	"fmt"
	"strconv"
)

func HandleDone(args []string, todos []todo.Todo) []todo.Todo {

	// Validation Layer
	if len(args) < 1 {
		fmt.Println(constants.MsgProvideTodoID)
		return todos
	}

	// Input Parsing
	id, err := strconv.Atoi(args[0])

	if err != nil {
		fmt.Println(constants.MsgInvalidTodoID)
		return todos
	}

	// Business Logic Layer
	todos, err = todo.MarkDone(todos, id)

	if err != nil {
		fmt.Println("Error:", err)
		return todos
	}

	// Persistence Layer
	err = storage.SaveTodos(todos)

    if err != nil {
	    fmt.Println("Error saving todos:", err)
	    return todos
    }

	return todos
}
