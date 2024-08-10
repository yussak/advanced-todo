package service

import (
	model "backend/models"
	repository "backend/repositories"
	"errors"
	"time"

	"github.com/oklog/ulid"
	"golang.org/x/exp/rand"
)

// todo:これrepo.todoでいいのか確認（model.Todoじゃなくていいのか）
func FetchTodos() ([]repository.Todo, error) {
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

	return req, nil
}
