package models

import "time"

// gorm.Model definition
type Model struct {
	ID        uint32 `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
