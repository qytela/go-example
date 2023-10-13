package main

import (
	"example/config"
	"example/routers"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	app := gin.Default()

	app.ForwardedByClientIP = true
	app.SetTrustedProxies([]string{"127.0.0.1"})

	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}

	// Initialize database
	config.DatabaseConnection()

	app.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello, mom!",
		})
	})

	// Set prefix
	apiRoutes := app.Group("/api")

	// Auth Routes
	routers.AuthRouter(apiRoutes.Group("/auth"))

	// User Routes
	routers.UserRouter(apiRoutes.Group("/users"))

	// Book Routes
	routers.BookRouter(apiRoutes.Group("/books"))

	app.Run()
}
