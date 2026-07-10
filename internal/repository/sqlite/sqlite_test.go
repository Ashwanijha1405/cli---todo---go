package sqlite

import (
	"os"
	"testing"

	"cli-todo/internal/repository/test"
)

func TestSQLiteRepository(t *testing.T) {
	db := "test_todos.db"

	// Cleanup before and after the test
	_ = os.Remove(db)
	defer os.Remove(db)

	repo, err := NewSQLiteRepository(db)
	if err != nil {
		t.Fatalf("failed to create SQLite repository: %v", err)
	}

	test.RunRepositoryTests(t, repo)
}