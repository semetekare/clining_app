package service

import (
	"clining_app/app/config"
	"clining_app/app/models"
)

type ClientService struct {
}

func NewClientService() *ClientService {
	return &ClientService{}
}

func (s *ClientService) CreateClient(client *models.Client) error {
	return config.DbPostgre.Create(client).Error
}

func (s *ClientService) GetClientByID(id uint) (*models.Client, error) {
	var client models.Client
	err := config.DbPostgre.First(&client, id).Error
	return &client, err
}

func (s *ClientService) GetAllClients() ([]models.Client, error) {
	var clients []models.Client
	err := config.DbPostgre.Find(&clients).Error
	return clients, err
}

func (s *ClientService) UpdateClient(client *models.Client) error {
	return config.DbPostgre.Save(client).Error
}

func (s *ClientService) DeleteClient(id uint) error {
	return config.DbPostgre.Delete(&models.Client{}, id).Error
}