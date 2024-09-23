package handlers

import (
	"net/http"

	"example.com/tuto/models"
	"github.com/gin-gonic/gin"
)

// GET /requirements
// Get all requirements
func FindRequirements(c *gin.Context) {
	var requirements []models.Requirement
	models.DB.Find(&requirements)

	c.JSON(http.StatusOK, gin.H{"data": requirements})
}

func CreateRequirement(c *gin.Context) {
	// Validate input
	var input models.CreateRequirementInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create requirement
	requirement := models.Requirement{Name: input.Name, Details: input.Details}
	if err := models.DB.Create(&requirement).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": requirement})
}

// GET /Requirements/:id
// Find a requirement
func FindRequirement(c *gin.Context) {
	var requirement models.Requirement

	if err := models.DB.Where("id = ?", c.Param("id")).First(&requirement).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Requirement not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": requirement})
}

// PATCH /requirements/:id
// Update a requirement
func UpdateRequirement(c *gin.Context) {
	// Get model if exist
	var Requirement models.Requirement
	if err := models.DB.Where("id = ?", c.Param("id")).First(&Requirement).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Requirement not found!"})
		return
	}

	// Validate input
	var input models.UpdateRequirementInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&Requirement).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": Requirement})
}

// DELETE /requirements/:id
// Delete a requirement
func DeleteRequirement(c *gin.Context) {
	// Get model if exist
	var requirement models.Requirement
	if err := models.DB.Where("id = ?", c.Param("id")).First(&requirement).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Requirement not found!"})
		return
	}

	models.DB.Delete(&requirement)

	c.Status(http.StatusNoContent)
}
