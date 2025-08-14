package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ClientID         uint
	Client           Client
	OrderDate        time.Time `gorm:"default:current_timestamp"`
	ServicePriceID   uint      // Assuming not null, adjust if needed
	ServicePrice     ServicePrice
	ExtraServiceID   *uint     // Assuming nullable if optional
	ExtraService     *ExtraService
}