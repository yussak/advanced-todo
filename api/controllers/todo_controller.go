package controller

import (
	"backend/internal/db"
	model "backend/models"
	repository "backend/repositories"
	service "backend/services"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO:共通化
type Todo struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

// TODO: fat controllerになっているので改善
// TODO:まずはcontroller内で適切に関数分離する

// コントローラにはリクエストを受け取って処理をしてviewに返すものを書く予定

func HandleFetchTodos(c *gin.Context) {
	todos, err := service.FetchTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todos)
}

func HandleAddTodo(c *gin.Context) {
	var req model.Todo
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo, err := service.PrepareTodo(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = repository.InsertTodoToDB(todo)
	if err != nil {
		// TODO:err.Error()としたら内部的なものが画面に表示されてしまうので治すかもしれない（他の部分も同じ）
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// c.JSON(http.StatusBadRequest, gin.H{"error": "Title and Body needed"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID must be provided"})
		return
	}

	_, err := db.DB.Exec("DELETE FROM todos WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func ShowTodo(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID must be provided"})
		return
	}

	row := db.DB.QueryRow("SELECT id, title, body FROM todos WHERE id = $1", id)

	var todo Todo
	err := row.Scan(&todo.ID, &todo.Title, &todo.Body)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "No todo with the provided ID."})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")

	var todo Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format or missing required fields"})
		return
	}

	sql := `UPDATE todos SET title = $1, body = $2 WHERE id = $3`
	if _, execErr := db.DB.Exec(sql, todo.Title, todo.Body, id); execErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": execErr.Error()})
		return
	}

	row := db.DB.QueryRow("SELECT id, title, body FROM todos WHERE id = $1", id)
	if err := row.Scan(&todo.ID, &todo.Title, &todo.Body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}
