package data

import (
	"encoding/json"
	"io"
	"time"
)

type Todo struct {
	ID        int
	Name      string
	Description string
	Deadline string
	StartedAt string
	isDone bool
	CreatedAt string
	Author string
}

type Todos []*Todo

func (todos Todos) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(todos)
}

func GetTodos() Todos {
	return TodosList
}

func AddTodo(todo *Todo) {
	todo.ID = getNextID()
	TodosList = append(TodosList, todo)
}

func getNextID() int {
	return TodosList[len(TodosList) - 1].ID + 1
}

var TodosList = Todos{
	&Todo{
		ID:        1,
		Name:      "Todo 1",
		StartedAt: time.Now().UTC().String(),
		Deadline:  time.Now().UTC().Add(time.Hour * time.Duration(5)).String(),
	},
	&Todo{
		ID:        2,
		Name:      "Todo 2",
		StartedAt: time.Now().UTC().String(),
		Deadline:  time.Now().UTC().Add(time.Hour * time.Duration(5)).String(),
		isDone: true,
	},
	&Todo{
		ID:        3,
		Name:      "Todo 3",
		StartedAt: time.Now().UTC().String(),
		Deadline:  time.Now().UTC().Add(time.Hour * time.Duration(5)).String(),
	},
}
