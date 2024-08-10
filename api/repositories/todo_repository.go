package repository

import (
	"backend/internal/db"
	model "backend/models"
	"database/sql"
	"errors"
)

// DBとのやり取りを担当

// TODO:代入、判定を一行で書き換えられる部分を書き換える→そうしたほうが見やすい場所だけやる

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

func UpdateTodoInDB(id string, todo model.Todo) error {
	sql := `UPDATE todos SET title = $1, body = $2 WHERE id = $3`
	if _, err := db.DB.Exec(sql, todo.Title, todo.Body, id); err != nil {
		return err
	}

	return nil
}
