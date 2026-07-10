package sqlite

import (
    "database/sql"
	"cli-todo/internal/repository/contract"
	 _ "modernc.org/sqlite"
	"errors"
	"cli-todo/internal/constants"
)

type SQLiteRepository struct {
    db *sql.DB
}

func NewSQLiteRepository(dbPath string) (*SQLiteRepository, error) {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}

	repo := &SQLiteRepository{
		db: db,
	}

	err = repo.init()

	if err != nil {
		return nil, err
	}

	return repo, nil
}

func (r *SQLiteRepository) init() error {
	query := `
	CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		completed BOOLEAN NOT NULL
	);`

	_, err := r.db.Exec(query)
	return err
}

func (r *SQLiteRepository) Create(title string) (contract.Todo, error) {
	query := `
	INSERT INTO todos(title, completed)
	VALUES(?, ?)
	`

	result, err := r.db.Exec(query, title, false)
	if err != nil {
		return contract.Todo{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return contract.Todo{}, err
	}

	return contract.Todo{
		ID:        int(id),
		Title:     title,
		Completed: false,
	}, nil
}

func (r *SQLiteRepository) GetAll() ([]contract.Todo, error) {
	query := `
	SELECT id, title, completed
	FROM todos
	ORDER BY id;
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []contract.Todo

	for rows.Next() {
		var todo contract.Todo

		err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed)
		if err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

func (r *SQLiteRepository) Update(todo contract.Todo) error {
	query := `
	UPDATE todos
	SET title = ?, completed = ?
	WHERE id = ?;
	`

	result, err := r.db.Exec(query, todo.Title, todo.Completed, todo.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New(constants.ErrTodoNotFound)
	}

	return nil
}

func (r *SQLiteRepository) Delete(id int) error {
	query := `
	DELETE FROM todos
	WHERE id = ?;
	`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New(constants.ErrTodoNotFound)
	}

	return nil
}


var _ contract.TodoRepository = (*SQLiteRepository)(nil)