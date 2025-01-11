package handlers

import (
	"log/slog"
	"net/http"

	"github.com/vbishel/trackio-backend/data"
)

type Todos struct {
	logger *slog.Logger
}

func NewTodos(logger *slog.Logger) *Todos {
	return &Todos{logger}
}

func (todos *Todos) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	todos.logger.Info("Handle GET Todos")

	todoList := data.GetTodos()

	err := todoList.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Failed to convert todolist to JSON", http.StatusInternalServerError)
		return
	}
}