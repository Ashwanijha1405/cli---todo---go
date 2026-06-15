package todo

import (
	"cli-todo/internal/constants"
	"cli-todo/internal/logger"
	"cli-todo/internal/repository"
	"errors"
	"fmt"
)

type TodoService struct {
	repo repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) ListTodos() ([]repository.Todo, error) {
	todos, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	logger.Log.Info("listing todos", "count", len(todos))

	for _, todo := range todos {
		fmt.Println(todo.ID, todo.Title, todo.Completed)
	}

	return todos, nil
}

func (s *TodoService) AddTodo(title string) (repository.Todo, error) {
	newTodo, err := s.repo.Create(title)
	if err != nil {
		return repository.Todo{}, err
	}

	logger.Log.Info(
		"todo created",
		"id", newTodo.ID,
		"title", newTodo.Title,
	)

	fmt.Println("Todo added:", title)

	return newTodo, nil
}

func (s *TodoService) MarkDone(id int) error {
	todos, err := s.repo.GetAll()
	if err != nil {
		return err
	}

	var foundTodo *repository.Todo
	for i := range todos {
		if todos[i].ID == id {
			todos[i].Completed = true
			foundTodo = &todos[i]
			break
		}
	}

	if foundTodo == nil {
		logger.Log.Error(
			"todo not found for completion",
			"id", id,
		)
		return errors.New(constants.ErrTodoNotFound)
	}

	err = s.repo.Update(*foundTodo)
	if err != nil {
		return err
	}

	logger.Log.Info(
		"todo marked completed",
		"id", id,
	)

	return nil
}

func (s *TodoService) DeleteTodo(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		logger.Log.Error(
			"todo not found for deletion",
			"id", id,
		)
		return err
	}

	logger.Log.Info(
		"todo deleted",
		"id", id,
	)

	return nil
}