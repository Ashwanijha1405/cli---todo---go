package repository

import (
	"fmt"
    "cli-todo/internal/config"
    contract "cli-todo/internal/repository/contract"
    jsonrepo "cli-todo/internal/repository/json"
    "cli-todo/internal/repository/sqlite"
)

func NewRepository(cfg config.Config) (contract.TodoRepository, error) {

	switch cfg.StorageType {

	case "json":
		return jsonrepo.NewJSONRepository("todos.json"), nil

	case "sqlite":
		return sqlite.NewSQLiteRepository(cfg.DBPath)

	default:
		return nil, fmt.Errorf("unsupported storage type: %s", cfg.StorageType)
	}
}