package main

import (
	"cli-todo/internal/todo"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

var todos []todo.Todo

func saveTodos() {

	data, err := json.MarshalIndent(todos, "", " ")

	if err != nil {
		fmt.Println("Error converting todos to json:", err)
		return
	}

	err = os.WriteFile("todos.json", data, 0644)

	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
}

func loadTodos() {

	data, err := os.ReadFile("todos.json")

	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	err = json.Unmarshal(data, &todos)

	if err != nil {
		fmt.Println("Error decoding file:", err)
		return
	}
}

func main() {

	loadTodos()

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

		saveTodos()

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

		saveTodos()

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

		saveTodos()

	default:
		fmt.Println("Unknown command")
	}
}
