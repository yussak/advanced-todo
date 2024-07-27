package router

import (
	controller "backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	// TODO:main.goなどでcors読み込むように変える
	r.GET("/todos", func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		controller.FetchTodos(c)
	})
}
