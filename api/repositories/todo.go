package repository

import (
	"backend/internal/db"
	model "backend/models"
)

// TODO:共通化
type Todo struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

// TODO:model.Todoに変える
func FetchTodosFromDB() ([]Todo, error) {
	rows, err := db.DB.Query("SELECT * FROM todos")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	todos := []Todo{}
	for rows.Next() {
		var todo Todo
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
