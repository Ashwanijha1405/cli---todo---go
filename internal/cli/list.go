package cli

import (
	"cli-todo/internal/todo"
	"fmt"
)

func HandleList(todoService *todo.TodoService) {

	_, err := todoService.ListTodos()
	if err != nil {
		fmt.Println("Error listing todos:", err)
	}
}
