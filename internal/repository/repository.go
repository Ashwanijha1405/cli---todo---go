package repository

type Todo struct {
	ID        int    `json:"ID"`
	Title     string `json:"Title"`
	Completed bool   `json:"Completed"`
}

type TodoRepository interface {
	Create(title string) (Todo, error)
	GetAll() ([]Todo, error)
	Update(todo Todo) error
	Delete(id int) error
}
