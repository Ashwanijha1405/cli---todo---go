package storage

import (
	"cli-todo/internal/logger"
	"cli-todo/internal/constants"
	"cli-todo/internal/todo"
	"encoding/json"
	"os"
)

func SaveTodos(todos []todo.Todo) error {

	logger.Log.Info("saving todos", "count", len(todos))

	data, err := json.MarshalIndent(todos, "", " ")

	if err != nil {
		logger.Log.Error("failed to convert todos to json", "error", err)
		return err
	}

	err = os.WriteFile(constants.TodosFile, data, 0644)

	if err != nil {
		logger.Log.Error("failed to write todos file", "error", err)
		return err
	}

	logger.Log.Info("todos saved successfully")

	return nil
}

func LoadTodos() ([]todo.Todo, error) {

	logger.Log.Info("loading todos", "file", constants.TodosFile)

	var todos []todo.Todo

	data, err := os.ReadFile(constants.TodosFile)

	if err != nil {

		if os.IsNotExist(err) {
			logger.Log.Info("todos file does not exist, creating empty todo list")
			return todos, nil
		}

		logger.Log.Error("failed to read todos file", "error", err)
		return todos, err
	}

	err = json.Unmarshal(data, &todos)

	if err != nil {
		logger.Log.Error("failed to decode todos file", "error", err)
		return todos, err
	}

	logger.Log.Info("todos loaded successfully", "count", len(todos))

	return todos, nil
}