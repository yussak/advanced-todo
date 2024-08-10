package controller

import (
	"backend/internal/db"
	service "backend/services"
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid"
	"golang.org/x/exp/rand"
)

// TODO:共通化
type Todo struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

// TODO: fat controllerになっているので改善
// TODO:まずはcontroller内で適切に関数分離する

func HandleFetchTodos(c *gin.Context) {
	todos, err := service.FetchTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todos)
}

func AddTodo(c *gin.Context) {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(uint64(t.UnixNano()))), 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy)

	var req Todo
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Title == "" || req.Body == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title and Body needed"})
		return
	}

	req.ID = id.String()
	sql := `INSERT INTO todos (id, title, body) VALUES($1, $2, $3)`
	_, err := db.DB.Exec(sql, req.ID, req.Title, req.Body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, req)
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
