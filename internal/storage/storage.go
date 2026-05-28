package storage

import (
	"cli-todo/internal/todo"
	"encoding/json"
	"os"
)

func SaveTodos(todos []todo.Todo) error {

	data, err := json.MarshalIndent(todos, "", " ")

	if err != nil {
		return err
	}

	err = os.WriteFile("todos.json", data, 0644)

	if err != nil {
		return err
	}

	return nil
}

func LoadTodos() ([]todo.Todo, error) {

	var todos []todo.Todo

	data, err := os.ReadFile("todos.json")

	if err != nil {

		if os.IsNotExist(err) {
			return todos, nil
		}

		return todos, err
	}

	err = json.Unmarshal(data, &todos)

	if err != nil {
		return todos, err
	}

	return todos, nil
}