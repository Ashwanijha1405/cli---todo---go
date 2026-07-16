package http

import (
	"net/http"

	"cli-todo/internal/http/handlers"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.HealthHandler)

	return mux
}