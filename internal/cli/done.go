package cli

import (
	"cli-todo/internal/constants"
	"cli-todo/internal/todo"
	"fmt"
	"strconv"
)

func HandleDone(args []string, todoService *todo.TodoService) {

	// Validation Layer
	if len(args) < 1 {
		fmt.Println(constants.MsgProvideTodoID)
		return
	}

	// Input Parsing
	id, err := strconv.Atoi(args[0])

	if err != nil {
		fmt.Println(constants.MsgInvalidTodoID)
		return
	}

	// Business Logic Layer & Persistence (via Service)
	err = todoService.MarkDone(id)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
