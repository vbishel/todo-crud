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

	if req.Method == http.MethodGet {
		todos.getProducts(rw, req)
	}

	if req.Method == http.MethodPost {
		todos.postProduct(rw, req)
	}
}

func (todos Todos) getProducts(rw http.ResponseWriter, _ *http.Request) {
	todos.logger.Info("Handling GET products")

	todoList := data.GetTodos()

	err := todoList.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Failed to convert todolist to JSON", http.StatusInternalServerError)
		todos.logger.Warn("[GET TODO]: Failed to convert to JSON")
		return
	}
}

func (todos Todos) postProduct(rw http.ResponseWriter, req *http.Request) {
	todos.logger.Info("Handling POST products")

	todo := &data.Todo{}
	
	err := todo.FromJSON(req.Body)
	if err != nil {
		todos.logger.Info(err.Error())
		http.Error(rw, "Failed to convert to JSON", http.StatusBadRequest)
		todos.logger.Warn("[POST TODO]: Failed to convert to JSON")
		return
	}

	data.AddTodo(todo)
}
