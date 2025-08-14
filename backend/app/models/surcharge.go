package models

import (
	"time"

	"gorm.io/gorm"
)

type Surcharge struct {
	gorm.Model
	Name       string    `gorm:"type:varchar(100);unique;not null"`
	Multiplier float64   `gorm:"type:numeric(3,2);not null"`
	StartTime  *time.Time // Nullable, using time.Time for TIME type
	EndTime    *time.Time // Nullable, using time.Time for TIME type
}