package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"example.com/tuto/config"
	"example.com/tuto/handlers"
	"example.com/tuto/middlewares"
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
	router.Use(middlewares.TrailingSlashMiddleware())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	handlers.SetupRoutes(router)

	router.Run("localhost:5000")
}
