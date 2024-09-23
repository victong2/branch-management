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

// POST /branches
// Create new branches
func CreateBranch(c *gin.Context) {
	// Validate input
	var input models.CreateBranchInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Make sure there is only one root branch.
	if input.ParentID == nil {
		var count int64
		if err := models.DB.Model(&models.Branch{}).Where("parent_id IS NULL").Count(&count).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Root branch already exists. Cannot create a branch with no parent."})
			return
		}
	}

	// Create branch
	branch := models.Branch{Name: input.Name}
	if input.ParentID != nil {
		branch.ParentID = input.ParentID
	}
	if err := models.DB.Create(&branch).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": branch})
}

// GET /Branches/:id
// Find a branch
func FindBranch(c *gin.Context) {
	var branch models.Branch

	if err := models.DB.Where("id = ?", c.Param("id")).First(&branch).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": branch})
}

// PATCH /branches/:id
// Update a branch
func UpdateBranch(c *gin.Context) {
	// Get model if exist
	var branch models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&branch).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch not found!"})
		return
	}

	// Validate input
	var input models.UpdateBranchInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&branch).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": branch})
}
