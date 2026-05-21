package main

import (
	"cli-todo/internal/storage"
	"cli-todo/internal/todo"
	"fmt"
	"os"
	"strconv"
)

func main() {

	todos := storage.LoadTodos()

	if len(os.Args) < 2 {
		fmt.Println("Please provide a command")
		return
	}

	command := os.Args[1]

	switch command {

	case "list":

		todo.ListTodos(todos)

	case "add":

		if len(os.Args) < 3 {
			fmt.Println("Please provide todo title")
			return
		}

		title := os.Args[2]

		todos = todo.AddTodo(todos, title)

		storage.SaveTodos(todos)

		todo.ListTodos(todos)

	case "done":

		if len(os.Args) < 3 {
			fmt.Println("Please provide Todo ID")
			return
		}

		id, err := strconv.Atoi(os.Args[2])

		if err != nil {
			fmt.Println("Invalid ID")
			return
		}

		todos = todo.MarkDone(todos, id)

		storage.SaveTodos(todos)

	case "delete":

		if len(os.Args) < 3 {
			fmt.Println("Please provide Todo ID")
			return
		}

		id, err := strconv.Atoi(os.Args[2])

		if err != nil {
			fmt.Println("Invalid ID")
			return
		}

		todos = todo.DeleteTodo(todos, id)

		storage.SaveTodos(todos)

	default:
		fmt.Println("Unknown command")
	}
}
