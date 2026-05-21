package storage

import (
	"cli-todo/internal/todo"
	"encoding/json"
	"fmt"
	"os"
)

func SaveTodos(todos []todo.Todo) {

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

func LoadTodos() []todo.Todo {

	var todos []todo.Todo

	data, err := os.ReadFile("todos.json")

	if err != nil {
		fmt.Println("Error reading file:", err)
		return todos
	}

	err = json.Unmarshal(data, &todos)

	if err != nil {
		fmt.Println("Error decoding file:", err)
		return todos
	}

	return todos
}
