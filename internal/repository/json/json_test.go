package json

import (
	"os"
	"testing"

	"cli-todo/internal/repository/test"
)

func TestJSONRepository(t *testing.T) {
	file := "test_todos.json"

	// Cleanup before and after the test
	_ = os.Remove(file)
	defer os.Remove(file)

	repo := NewJSONRepository(file)

	test.RunRepositoryTests(t, repo)
}