package controller

import (
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

// TODO:代入、判定を位置行で書き換えられる部分を書き換える（そこでしか使ってない変数は行けるはず）

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

func HandleUpdateTodo(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID must be provided"})
		return
	}

	var todo model.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format or missing required fields"})
		return
	}

	todo, err := service.UpdateTodo(id, todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}
