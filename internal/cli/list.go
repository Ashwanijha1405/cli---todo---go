package cli

import "cli-todo/internal/todo"

func HandleList(todos []todo.Todo) []todo.Todo {

	todo.ListTodos(todos)

	return todos
}
