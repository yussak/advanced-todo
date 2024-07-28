package controller

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid"
	"golang.org/x/exp/rand"
)

type Todo struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

var db *sql.DB

// TODO:init db処理を切り出して読み込むように変える
func SetDB(database *sql.DB) {
	db = database
}

func FetchTodos(c *gin.Context) {
	rows, err := db.Query("SELECT * FROM todos")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer rows.Close()
	todos := []Todo{}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		todos = append(todos, todo)
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

	req.ID = id.String()
	sql := `INSERT INTO todos (id, title, body) VALUES($1, $2, $3)`
	_, err := db.Exec(sql, req.ID, req.Title, req.Body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, req)
}
