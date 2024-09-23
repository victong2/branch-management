package handlers

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		api.GET("/books", FindBooks)
		api.POST("/books", CreateBook)
		api.GET("/books/:id", FindBook)
		api.PATCH("/books/:id", UpdateBook)
		api.DELETE("/books/:id", DeleteBook)

		api.GET("/branches", FindBranches)
		api.POST("/branches", CreateBranch)
		api.GET("/branches/:id", FindBranch)
		api.PATCH("/branches/:id", UpdateBranch)
	}
}
