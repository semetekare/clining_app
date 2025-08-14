package models

import (
	"gorm.io/gorm"
)

type ExtraService struct {
	gorm.Model
	Name  string  `gorm:"type:varchar(200);unique;not null"`
	Unit  string  `gorm:"type:varchar(20);not null"`
	Price float64 `gorm:"type:numeric(10,2);not null"`
}