package main

import (
	"backend/internal/db"
	middleware "backend/middlewares"
	router "backend/routers"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	_, err := db.InitDB()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	r := gin.Default()
	r.Use(cors.New(middleware.CorsConfig()))
	router.SetupRouter(r)

	r.Run(":8080")
}
