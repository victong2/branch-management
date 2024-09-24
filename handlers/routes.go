package handlers

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		api.GET("/branches", FindBranches)
		api.POST("/branches", CreateBranch)
		api.GET("/branches/:id", FindBranch)
		api.PATCH("/branches/:id", UpdateBranch)
		api.DELETE("/branches/:id", DeleteBranch)

		api.GET("/requirements", FindRequirements)
		api.POST("/requirements", CreateRequirement)
		api.GET("/requirements/:id", FindRequirement)
		api.PATCH("/requirements/:id", UpdateRequirement)
		api.DELETE("/requirements/:id", DeleteRequirement)

		api.GET("/branches/:id/requirements", FindRequirementsByBranch)
		api.POST("/branches/:id/requirements", AddRequirements)
		api.PUT("/branches/:id/requirements", ReplaceRequirements)
	}
}
