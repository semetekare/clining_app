package service

import (
	"clining_app/app/config"
	"clining_app/app/models"
)
type ServiceService struct {
}

func NewServiceService() *ServiceService {
	return &ServiceService{}
}

func (s *ServiceService) CreateService(service *models.Service) error {
	return config.DbPostgre.Create(service).Error
}

func (s *ServiceService) GetServiceByID(id uint) (*models.Service, error) {
	var service models.Service
	err := config.DbPostgre.First(&service, id).Error
	return &service, err
}

func (s *ServiceService) GetAllServices() ([]models.Service, error) {
	var services []models.Service
	err := config.DbPostgre.Find(&services).Error
	return services, err
}

func (s *ServiceService) UpdateService(service *models.Service) error {
	return config.DbPostgre.Save(service).Error
}

func (s *ServiceService) DeleteService(id uint) error {
	return config.DbPostgre.Delete(&models.Service{}, id).Error
}