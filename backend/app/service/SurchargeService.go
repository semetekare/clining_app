package service

import (
	"clining_app/app/config"
	"clining_app/app/models"
)

type SurchargeService struct {
}

func NewSurchargeService() *SurchargeService {
	return &SurchargeService{}
}

func (s *SurchargeService) CreateSurcharge(surcharge *models.Surcharge) error {
	return config.DbPostgre.Create(surcharge).Error
}

func (s *SurchargeService) GetSurchargeByID(id uint) (*models.Surcharge, error) {
	var surcharge models.Surcharge
	err := config.DbPostgre.First(&surcharge, id).Error
	return &surcharge, err
}

func (s *SurchargeService) GetAllSurcharges() ([]models.Surcharge, error) {
	var surcharges []models.Surcharge
	err := config.DbPostgre.Find(&surcharges).Error
	return surcharges, err
}

func (s *SurchargeService) UpdateSurcharge(surcharge *models.Surcharge) error {
	return config.DbPostgre.Save(surcharge).Error
}

func (s *SurchargeService) DeleteSurcharge(id uint) error {
	return config.DbPostgre.Delete(&models.Surcharge{}, id).Error
}