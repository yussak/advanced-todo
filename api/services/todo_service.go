package service

import (
	model "backend/models"
	repository "backend/repositories"
	"errors"
	"time"

	"github.com/oklog/ulid"
	"golang.org/x/exp/rand"
)

func FetchTodos() ([]model.Todo, error) {
	todos, err := repository.FetchTodosFromDB()
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func PrepareTodo(req model.Todo) (model.Todo, error) {
	if req.Title == "" || req.Body == "" {
		return model.Todo{}, errors.New("title and body needed")
	}

	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(uint64(t.UnixNano()))), 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy)

	req.ID = id.String()

	err := repository.InsertTodoToDB(req)
	if err != nil {
		return model.Todo{}, err
	}

	return req, nil
}

func DeleteTodo(id string) error {
	err := repository.DeleteTodoFromDB(id)
	if err != nil {
		return err
	}

	return nil
}
