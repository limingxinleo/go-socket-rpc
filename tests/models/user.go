package models

import (
	"time"
)

type User struct {
	Model
	Name      string `gorm:"type:varchar(255);unique_index"`
	RoleId    int32  `gorm:"index"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
