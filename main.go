package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"example.com/tuto/config"
	"example.com/tuto/controllers"
	"example.com/tuto/models"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	models.ConnectDatabase(config)

	router := gin.Default()

	setupRoutes(router)

	router.Run("localhost:5000")
}

func setupRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	api := router.Group("/api/v1")
	{
		api.GET("/books", controllers.FindBooks)
		api.POST("/books", controllers.CreateBook)
		api.GET("/books/:id", controllers.FindBook)
		api.PATCH("/books/:id", controllers.UpdateBook)
		api.DELETE("/books/:id", controllers.DeleteBook)
	}
}
