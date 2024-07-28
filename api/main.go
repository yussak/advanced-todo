package main

import (
	controller "backend/controllers"
	middleware "backend/middlewares"
	router "backend/routers"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DBNAME"))

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connection established!")

	// TODO:init db処理を切り出して読み込むように変える
	controller.SetDB(db)

	r := gin.Default()
	r.Use(cors.New(middleware.CorsConfig()))
	router.SetupRouter(r)

	r.Run(":8080")
}
