package models

type Branch struct {
	ID           uint          `gorm:"primary_key" json:"id"`
	Name         string        `gorm:"type:varchar(100);not null" json:"name"`
	ParentID     *uint         `gorm:"index" json:"parent_id"` // Nullable for top-level branches
	Requirements []Requirement `gorm:"many2many:branch_requirements;" json:"requirements"`
}

type CreateBranchInput struct {
	Name     string `json:"name" binding:"required"`
	ParentID *uint  `json:"parent_id,omitempty"`
}

type UpdateBranchInput struct {
	Name     string `json:"name"`
	ParentID *uint  `json:"parent_id,omitempty"`
}
