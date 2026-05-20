package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos []Todo

func listTodos() {
	for _, todo := range todos {
		fmt.Println(todo.ID, todo.Title, todo.Completed)
	}
}

func addTodo(title string) {
	newTodo := Todo{
		ID:        len(todos) + 1,
		Title:     title,
		Completed: false,
	}

	todos = append(todos, newTodo)

	saveTodos()

	fmt.Println("Todo added: ", title)
}

func saveTodos() {

	data, err := json.MarshalIndent(todos, "", " ")

	if err != nil {
		fmt.Println("Error converting todos to json:", err)
		return
	}

	err = os.WriteFile("todos.json", data, 0644)

	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
}

func loadTodos() {

	data, err := os.ReadFile("todos.json")

	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	err = json.Unmarshal(data, &todos)

	if err != nil {
		fmt.Println("Error decoding file:", err)
		return
	}
}

func markDone(id int) {

	for i, todo := range todos {

		if todo.ID == id {

			todos[i].Completed = true

			saveTodos()

			fmt.Println("Todo marked as completed!")
			return
		}
	}

	fmt.Println("Todo not found")
}

func deleteTodo(id int) {

	for i, todo := range todos {

		if todo.ID == id {

			todos = append(todos[:i], todos[i+1:]...)

			saveTodos()

			fmt.Println("Todos Deleted!")
			return
		}

	}

	fmt.Println("Todo not found.")

}

func main() {

	loadTodos()

	command := os.Args[1]

	switch command {

	case "list":
		listTodos()

	case "add":

		if len(os.Args) < 3 {
			fmt.Println("Please provide todo title")
			return
		}

		title := os.Args[2]

		addTodo(title)

		listTodos()

	case "done":

		if len(os.Args) < 3 {
			fmt.Println("Please Provide Todo ID")
			return
		}

		id, err := strconv.Atoi(os.Args[2])

		if err != nil {
			fmt.Println("Invalid ID")
			return
		}

		markDone(id)

	case "delete":

		if len(os.Args) < 3 {
			fmt.Println("Please provide Todo ID.")
			return
		}

		id, err := strconv.Atoi(os.Args[2])

		if err != nil {
			fmt.Println("Invalid ID")
			return
		}

		deleteTodo(id)

	default:
		fmt.Println("Unknown command")
	}
}
