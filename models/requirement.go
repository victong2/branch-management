package models

// Requirement model
type Requirement struct {
	ID       uint     `gorm:"primaryKey" json:"id"`
	Name     string   `gorm:"type:varchar(100);not null" json:"name"`
	Details  string   `gorm:"type:text;" json:"details"`
	Branches []Branch `gorm:"many2many:branch_requirements;" json:"branches"`
}
