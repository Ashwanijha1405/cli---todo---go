package main

import (
	"log"
	"net/http"

	apphttp "cli-todo/internal/http"
)

func main() {
	router := apphttp.NewRouter()

	log.Println("Server listening on :8080")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}