package service

import (
	model "backend/models"
	repository "backend/repositories"
	"errors"
	"time"

	"github.com/oklog/ulid"
	"golang.org/x/exp/rand"
)

// TODO:代入、判定を一行で書き換えられる部分を書き換える→そうしたほうが見やすい場所だけやる

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

	req.ID = generateTodoID()

	err := repository.InsertTodoToDB(req)
	if err != nil {
		return model.Todo{}, err
	}

	return req, nil
}

func generateTodoID() string {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(uint64(t.UnixNano()))), 0)
	ID := ulid.MustNew(ulid.Timestamp(t), entropy)

	// stringに変換したidを返す
	return ID.String()
}

func DeleteTodo(id string) error {
	err := repository.DeleteTodoFromDB(id)
	if err != nil {
		return err
	}

	return nil
}

func ShowTodo(id string) (model.Todo, error) {

	todo, err := repository.FetchTodoDetailFromDB(id)
	if err != nil {
		return model.Todo{}, err
	}

	return todo, nil
}

func UpdateTodo(id string, todo model.Todo) (model.Todo, error) {
	err := repository.UpdateTodoInDB(id, todo)
	if err != nil {
		return model.Todo{}, err
	}

	return todo, nil
}
