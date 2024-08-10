package router

import (
	controller "backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.GET("/todos", controller.HandleFetchTodos)
	r.POST("/todo", controller.HandleAddTodo)
	r.DELETE("/todo/:id", controller.HandleDeleteTodo)
	r.GET("/todos/:id", controller.ShowTodo)
	r.PUT("/todos/edit/:id", controller.UpdateTodo)
}
