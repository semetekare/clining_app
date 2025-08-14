package service

import (
	"clining_app/app/config"
	"clining_app/app/models"
)

type RoomTypeService struct {
}

func NewRoomTypeService() *RoomTypeService {
	return &RoomTypeService{}
}

func (s *RoomTypeService) CreateRoomType(roomType *models.RoomType) error {
	return config.DbPostgre.Create(roomType).Error
}

func (s *RoomTypeService) GetRoomTypeByID(id uint) (*models.RoomType, error) {
	var roomType models.RoomType
	err := config.DbPostgre.First(&roomType, id).Error
	return &roomType, err
}

func (s *RoomTypeService) GetAllRoomTypes() ([]models.RoomType, error) {
	var roomTypes []models.RoomType
	err := config.DbPostgre.Find(&roomTypes).Error
	return roomTypes, err
}

func (s *RoomTypeService) UpdateRoomType(roomType *models.RoomType) error {
	return config.DbPostgre.Save(roomType).Error
}

func (s *RoomTypeService) DeleteRoomType(id uint) error {
	return config.DbPostgre.Delete(&models.RoomType{}, id).Error
}