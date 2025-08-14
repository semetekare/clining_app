package models

import (
	"gorm.io/gorm"
)

type ServicePrice struct {
	gorm.Model
	RoomTypeID    uint
	RoomType      RoomType
	BasePrice     float64 `gorm:"type:numeric(10,2);not null"`
	BaseDuration  float64 `gorm:"type:numeric(4,1)"` // Nullable? Assuming optional but based on SQL not specified as NULL
	OvertimeRate  float64 `gorm:"type:numeric(10,2)"` // Nullable? Assuming optional
	Specialists   int16   `gorm:"not null"`           // SMALLINT
	SurchargeID   *uint   // Nullable? Assuming optional based on comment
	Surcharge     *Surcharge
}