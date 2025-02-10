package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint           `gorm:"primaryKey"`
	Name      string         `json:"name"`
	Age       int            `json:"age"`
	Address   string         `json:"address"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
