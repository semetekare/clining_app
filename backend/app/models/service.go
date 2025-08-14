package models

import (
	"gorm.io/gorm"
)

type Service struct {
	gorm.Model
	Name string `gorm:"type:varchar(100);unique;not null"`
}