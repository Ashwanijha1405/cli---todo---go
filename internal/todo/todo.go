package todo

import (
	"cli-todo/internal/logger"
	"cli-todo/internal/constants"
	"errors"
	"fmt"
)

type Todo struct {
	ID        int
	Title     string
	Completed bool
}

func ListTodos(todos []Todo) {

	logger.Log.Info("listing todos", "count", len(todos))

	for _, todo := range todos {
		fmt.Println(todo.ID, todo.Title, todo.Completed)
	}
}

func AddTodo(todos []Todo, title string) []Todo {

	newTodo := Todo{
		ID:        getNextID(todos),
		Title:     title,
		Completed: false,
	}

	todos = append(todos, newTodo)

	logger.Log.Info(
		"todo created",
		"id", newTodo.ID,
		"title", newTodo.Title,
	)

	fmt.Println("Todo added:", title)

	return todos
}

func MarkDone(todos []Todo, id int) ([]Todo, error) {

	for i, todo := range todos {

		if todo.ID == id {

			todos[i].Completed = true

			logger.Log.Info(
				"todo marked completed",
				"id", id,
			)

			return todos, nil
		}
	}

	logger.Log.Error(
		"todo not found for completion",
		"id", id,
	)

	return todos, errors.New(constants.ErrTodoNotFound)
}

func DeleteTodo(todos []Todo, id int) ([]Todo, error) {

	for i, todo := range todos {

		if todo.ID == id {

			todos = append(todos[:i], todos[i+1:]...)

			logger.Log.Info(
				"todo deleted",
				"id", id,
			)

			return todos, nil
		}
	}

	logger.Log.Error(
		"todo not found for deletion",
		"id", id,
	)

	return todos, errors.New(constants.ErrTodoNotFound)
}

func getNextID(todos []Todo) int {

	maxID := 0

	for _, todo := range todos {

		if todo.ID > maxID {
			maxID = todo.ID
		}
	}

	return maxID + 1
}