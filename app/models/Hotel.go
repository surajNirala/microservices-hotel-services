package models

import (
	"time"

	"gorm.io/gorm"
)

type Hotel struct {
	gorm.Model
	ID        uint       `gorm:"primaryKey;column:id" json:"id"`
	Name      string     `gorm:"size:256;column:name" json:"name" binding:"required"`
	UserID    uint       `gorm:"size:10;" json:"user_id"`
	Status    bool       `gorm:"default:true" json:"status"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}
