package handlers

import (
	"net/http"

	"example.com/tuto/models"
	"github.com/gin-gonic/gin"
)

// GET /branches/:branchID/requirements
// Find requirements for a branch
func FindRequirementsByBranch(c *gin.Context) {
	branchID := c.Param("id")

	var branch models.Branch
	if err := models.DB.First(&branch, branchID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Branch not found: " + err.Error()})
		return
	}

	// Merge branch requirements with parents/ancestors requirements.
	requirements := make(map[uint]models.Requirement)
	for id := &branch.ID; id != nil; {
		var br models.Branch
		if err := models.DB.Preload("Requirements").First(&br, *id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Branch not found: " + err.Error()})
			return
		}
		for _, r := range br.Requirements {
			_, ok := requirements[r.ID]
			if !ok {
				requirements[r.ID] = r
			}
		}
		id = br.ParentID
	}

	// Convert requirements map to slice
	branch.Requirements = make([]models.Requirement, 0, len(requirements))
	for _, r := range requirements {
		branch.Requirements = append(branch.Requirements, r)
	}

	c.JSON(http.StatusOK, gin.H{"data": branch})
}

// TODO: extract common code between AddRequirements and ReplaceRequirements.

// POST /branches/:branchID/requirements
// Add requirements to a branch
func AddRequirements(c *gin.Context) {
	branchID := c.Param("id")

	var input models.AssignRequirementsInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retrieve the branch from the database.
	// Fetch requirements for that branch.
	var branch models.Branch
	if err := models.DB.First(&branch, branchID).Error; err != nil {
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

// PATCH /branches/:branchID/requirements
// Replace requirements for a branch
func ReplaceRequirements(c *gin.Context) {
	branchID := c.Param("id")

	var input models.AssignRequirementsInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retrieve the branch from the database.
	// Fetch requirements for that branch.
	var branch models.Branch
	if err := models.DB.First(&branch, branchID).Error; err != nil {
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
