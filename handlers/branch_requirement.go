package handlers

import (
	"net/http"

	"example.com/tuto/models"
	"github.com/gin-gonic/gin"
)

// TODO: extract common code between AddRequirements and ReplaceRequirements.

func AddRequirements(c *gin.Context) {
	branchID := c.Param("branchID")

	var input models.AssignRequirementsInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retrieve the branch from the database.
	// Fetch requirements for that branch.
	var branch models.Branch
	if err := models.DB.Preload("Requirements").First(&branch, branchID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Branch not found"})
		return
	}

	// Fetch and assign the requirements to the branch.
	// Perform insertions in the junction table.
	var requirements []models.Requirement
	models.DB.Where("id IN ?", input.Requirements).Find(&requirements)
	models.DB.Model(&branch).Association("Requirements").Append(requirements)

	// Return the updated branch with requirements
	models.DB.Preload("Requirements").First(&branch, branchID)
	c.JSON(http.StatusOK, branch)
}

func ReplaceRequirements(c *gin.Context) {
	branchID := c.Param("branchID")

	var input models.AssignRequirementsInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retrieve the branch from the database.
	// Fetch requirements for that branch.
	var branch models.Branch
	if err := models.DB.Preload("Requirements").First(&branch, branchID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Branch not found"})
		return
	}

	// Fetch and assign the requirements to the branch.
	// Perform insertions in the junction table.
	var requirements []models.Requirement
	models.DB.Where("id IN ?", input.Requirements).Find(&requirements)
	models.DB.Model(&branch).Association("Requirements").Replace(requirements)

	// Return the updated branch with requirements
	models.DB.Preload("Requirements").First(&branch, branchID)
	c.JSON(http.StatusOK, branch)
}
