package handlers

import (
	"net/http"

	"example.com/tuto/models"
	"github.com/gin-gonic/gin"
)

// GET /branches
// Get all branches
func FindBranches(c *gin.Context) {
	var branches []models.Branch
	models.DB.Find(&branches)

	c.JSON(http.StatusOK, gin.H{"data": branches})
}
