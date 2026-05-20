package todo

import "fmt"

type Todo struct {
	ID        int
	Title     string
	Completed bool
}

func ListTodos(todos []Todo) {
	for _, todo := range todos {
		fmt.Println(todo.ID, todo.Title, todo.Completed)
	}
}

func AddTodo(todos []Todo, title string) []Todo {
	newTodo := Todo{
		ID:        len(todos) + 1,
		Title:     title,
		Completed: false,
	}

	todos = append(todos, newTodo)

	fmt.Println("Todo added:", title)

	return todos
}

func MarkDone(todos []Todo, id int) []Todo {
	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Completed = true

			fmt.Println("Todo marked as completed!")
			return todos
		}
	}

	fmt.Println("Todo not found")
	return todos
}

func DeleteTodo(todos []Todo, id int) []Todo {
	for i, todo := range todos {
		if todo.ID == id {

			todos = append(todos[:i], todos[i+1:]...)

			fmt.Println("Todo deleted!")
			return todos
		}
	}

	fmt.Println("Todo not found.")
	return todos
}
