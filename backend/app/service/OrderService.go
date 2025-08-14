package service

import (
	"clining_app/app/config"
	"clining_app/app/models"
)
type OrderService struct {
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (s *OrderService) CreateOrder(order *models.Order) error {
	return config.DbPostgre.Create(order).Error
}

func (s *OrderService) GetOrderByID(id uint) (*models.Order, error) {
	var order models.Order
	err := config.DbPostgre.First(&order, id).Error
	return &order, err
}

func (s *OrderService) GetAllOrders() ([]models.Order, error) {
	var orders []models.Order
	err := config.DbPostgre.Find(&orders).Error
	return orders, err
}

func (s *OrderService) UpdateOrder(order *models.Order) error {
	return config.DbPostgre.Save(order).Error
}

func (s *OrderService) DeleteOrder(id uint) error {
	return config.DbPostgre.Delete(&models.Order{}, id).Error
}