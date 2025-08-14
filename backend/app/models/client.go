package models

import (
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	Name  string `gorm:"type:varchar(100);not null"`
	Phone string `gorm:"type:varchar(20);unique;not null"`
	Email string `gorm:"type:varchar(100)"` // Nullable
}
