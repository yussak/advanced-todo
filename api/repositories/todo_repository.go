package repository

import (
	"backend/internal/db"
	model "backend/models"
	"database/sql"
	"errors"
)

// DBとのやり取りを担当

func FetchTodosFromDB() ([]model.Todo, error) {
	rows, err := db.DB.Query("SELECT * FROM todos")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	todos := []model.Todo{}
	for rows.Next() {
		var todo model.Todo
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Body)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func InsertTodoToDB(todo model.Todo) error {
	sql := `INSERT INTO todos (id, title, body) VALUES($1, $2, $3)`
	_, err := db.DB.Exec(sql, todo.ID, todo.Title, todo.Body)
	if err != nil {
		return err
	}

	return nil
}

func DeleteTodoFromDB(id string) error {
	_, err := db.DB.Exec("DELETE FROM todos WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

func FetchTodoDetailFromDB(id string) (model.Todo, error) {
	row := db.DB.QueryRow("SELECT id, title, body FROM todos WHERE id = $1", id)

	var todo model.Todo
	err := row.Scan(&todo.ID, &todo.Title, &todo.Body)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Todo{}, errors.New("todo not found")
		}

		return model.Todo{}, nil
	}

	return todo, nil
}
