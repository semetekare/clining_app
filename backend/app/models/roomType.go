package models

import (
	"gorm.io/gorm"
)

type RoomType struct {
	gorm.Model
	ServiceID   uint
	Service     Service
	MinArea     *int // Nullable
	MaxArea     *int // Nullable
}