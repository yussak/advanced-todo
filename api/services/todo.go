package service

import (
	repository "backend/repositories"
)

func FetchTodos() ([]repository.Todo, error) {
	todos, err := repository.FetchTodosFromDB()
	if err != nil {
		return nil, err
	}

	return todos, nil
}
