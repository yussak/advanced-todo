package controller

import (
	"backend/internal/db"
	model "backend/models"
	service "backend/services"
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
// TODO:コントローラ層はリポジトリ層とやり取りをしないのに揃える

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
		// TODO:err.Error()としたら内部的なものが画面に表示されてしまうので治す（他の部分も同じ）→An error occurredのようにする
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func HandleDeleteTodo(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID must be provided"})
		return
	}

	err := service.DeleteTodo(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func HandleShowTodo(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID must be provided"})
		return
	}

	todo, err := service.ShowTodo(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
