package repository

import (
	"backend/internal/db"
	model "backend/models"
)

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
