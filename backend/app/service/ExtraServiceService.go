package service

import (
	"clining_app/app/config"
	"clining_app/app/models"
)

type ExtraServiceService struct {
}

func NewExtraServiceService() *ExtraServiceService {
	return &ExtraServiceService{}
}

func (s *ExtraServiceService) CreateExtraService(extraService *models.ExtraService) error {
	return config.DbPostgre.Create(extraService).Error
}

func (s *ExtraServiceService) GetExtraServiceByID(id uint) (*models.ExtraService, error) {
	var extraService models.ExtraService
	err := config.DbPostgre.First(&extraService, id).Error
	return &extraService, err
}

func (s *ExtraServiceService) GetAllExtraServices() ([]models.ExtraService, error) {
	var extraServices []models.ExtraService
	err := config.DbPostgre.Find(&extraServices).Error
	return extraServices, err
}

func (s *ExtraServiceService) UpdateExtraService(extraService *models.ExtraService) error {
	return config.DbPostgre.Save(extraService).Error
}

func (s *ExtraServiceService) DeleteExtraService(id uint) error {
	return config.DbPostgre.Delete(&models.ExtraService{}, id).Error
}