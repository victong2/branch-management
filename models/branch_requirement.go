package models

// BranchRequirement (junction table)
type BranchRequirement struct {
	BranchID      uint `gorm:"primaryKey" json:"branch_id"`
	RequirementID uint `gorm:"primaryKey" json:"requirement_id"`
}

type AssignRequirementsInput struct {
	Requirements []uint `json:"requirements" binding:"required"`
}
