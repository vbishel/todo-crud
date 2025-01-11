package data

import (
	"encoding/json"
	"io"
	"time"
)

type Todo struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Deadline    string `json:"deadline"`
	StartedAt   string `json:"startedAt"`
	IsDone      bool   `json:"isDone"`
	CreatedAt   string `json:"createdAt"`
	Author      string `json:"author"`
}

type Todos []*Todo

func (todos Todos) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(todos)
}

func (todo *Todo) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(todo)
}
 
func GetTodos() Todos {
	return TodosList
}

func AddTodo(todo *Todo) {
	todo.ID = getNextID()
	TodosList = append(TodosList, todo)
}

func getNextID() int {
	return TodosList[len(TodosList)-1].ID + 1
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
		IsDone:    true,
	},
	&Todo{
		ID:        3,
		Name:      "Todo 3",
		StartedAt: time.Now().UTC().String(),
		Deadline:  time.Now().UTC().Add(time.Hour * time.Duration(5)).String(),
	},
}
