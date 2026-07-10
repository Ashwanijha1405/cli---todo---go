package config

import (
	"errors"
	"os"
)

type Config struct {
	StorageType string
	DBPath      string
	LogLevel    string
}

func Load() (Config, error) {
	storageType := os.Getenv("STORAGE_TYPE")
	dbPath := os.Getenv("DB_PATH")
	logLevel := os.Getenv("LOG_LEVEL")

	if storageType == "" {
		storageType = "sqlite"
	}

	if dbPath == "" {
		dbPath = "todos.db"
	}

	if logLevel == "" {
		logLevel = "info"
	}

	if storageType != "sqlite" {
	return Config{}, errors.New("unsupported storage type: " + storageType)
    }

	cfg := Config{
		StorageType: storageType,
		DBPath:      dbPath,
		LogLevel:    logLevel,
	}

	return cfg, nil
}