package service

import (
	"clining_app/app/config"
	"clining_app/app/models"
)

type ServicePriceService struct {
}

func NewServicePriceService() *ServicePriceService {
	return &ServicePriceService{}
}

func (s *ServicePriceService) CreateServicePrice(servicePrice *models.ServicePrice) error {
	return config.DbPostgre.Create(servicePrice).Error
}

func (s *ServicePriceService) GetServicePriceByID(id uint) (*models.ServicePrice, error) {
	var servicePrice models.ServicePrice
	err := config.DbPostgre.First(&servicePrice, id).Error
	return &servicePrice, err
}

func (s *ServicePriceService) GetAllServicePrices() ([]models.ServicePrice, error) {
	var servicePrices []models.ServicePrice
	err := config.DbPostgre.Find(&servicePrices).Error
	return servicePrices, err
}

func (s *ServicePriceService) UpdateServicePrice(servicePrice *models.ServicePrice) error {
	return config.DbPostgre.Save(servicePrice).Error
}

func (s *ServicePriceService) DeleteServicePrice(id uint) error {
	return config.DbPostgre.Delete(&models.ServicePrice{}, id).Error
}