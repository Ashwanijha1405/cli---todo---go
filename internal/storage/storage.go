package storage

import (
	"cli-todo/internal/constants"
	"cli-todo/internal/logger"
	"cli-todo/internal/repository"
	"encoding/json"
	"errors"
	"os"
)

type JSONRepository struct {
	filePath string
}

func NewJSONRepository(filePath string) *JSONRepository {
	return &JSONRepository{filePath: filePath}
}

func (r *JSONRepository) GetAll() ([]repository.Todo, error) {
	logger.Log.Info("loading todos", "file", r.filePath)

	var todos []repository.Todo

	data, err := os.ReadFile(r.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			logger.Log.Info("todos file does not exist, creating empty todo list")
			return todos, nil
		}
		logger.Log.Error("failed to read todos file", "error", err)
		return nil, err
	}

	err = json.Unmarshal(data, &todos)
	if err != nil {
		logger.Log.Error("failed to decode todos file", "error", err)
		return nil, err
	}

	logger.Log.Info("todos loaded successfully", "count", len(todos))
	return todos, nil
}

func (r *JSONRepository) save(todos []repository.Todo) error {
	logger.Log.Info("saving todos", "count", len(todos))

	data, err := json.MarshalIndent(todos, "", " ")
	if err != nil {
		logger.Log.Error("failed to convert todos to json", "error", err)
		return err
	}

	err = os.WriteFile(r.filePath, data, 0644)
	if err != nil {
		logger.Log.Error("failed to write todos file", "error", err)
		return err
	}

	logger.Log.Info("todos saved successfully")
		logger.Log.Error("failed to write todos file", "error", err)
		return err
	}

	logger.Log.Info("todos saved successfully")

	return nil
}

func (r *JSONRepository) Create(title string) (repository.Todo, error) {
	todos, err := r.GetAll()
	if err != nil {
		return repository.Todo{}, err
	}

	maxID := 0
	for _, t := range todos {
		if t.ID > maxID {
			maxID = t.ID
		}
	}

	newTodo := repository.Todo{
		ID:        maxID + 1,
		Title:     title,
		Completed: false,
	}

	todos = append(todos, newTodo)
	err = r.save(todos)
	if err != nil {
		return repository.Todo{}, err
	}

	return newTodo, nil
}

func (r *JSONRepository) Update(updatedTodo repository.Todo) error {
	todos, err := r.GetAll()
	if err != nil {
		return err
	}

	found := false
	for i, t := range todos {
		if t.ID == updatedTodo.ID {
			todos[i] = updatedTodo
			found = true
			break
		}
	}

	if !found {
		return errors.New(constants.ErrTodoNotFound)
	}

	return r.save(todos)
}

func (r *JSONRepository) Delete(id int) error {
	todos, err := r.GetAll()
	if err != nil {
		return err
	}

	found := false
	var updatedTodos []repository.Todo
	for _, t := range todos {
		if t.ID == id {
			found = true
			continue
		}
		updatedTodos = append(updatedTodos, t)
	}

	if !found {
		return errors.New(constants.ErrTodoNotFound)
	}

	return r.save(updatedTodos)
}