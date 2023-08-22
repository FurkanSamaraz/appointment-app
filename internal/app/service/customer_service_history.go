package service

import (
	api_structure "meeting_app/internal/app/structures"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

//go:generate mockgen -destination=../../../tests/mocks/service/mock_customer_service_history.go -package=service meeting_app/internal/app/service ICustomerServiceHistory
type CustomerServiceHistoryService struct{ DB *gorm.DB }

type ICustomerServiceHistory interface {
	GetCustomerServiceHistory(filter api_structure.CustomerServiceHistoryFilter) ([]api_structure.CustomerServiceHistory, error)
	GetCustomerServiceHistoryById(id uuid.UUID) (api_structure.CustomerServiceHistory, error)
	UpdateCustomerServiceHistory(id uuid.UUID, data api_structure.CustomerServiceHistoryEdit) error
	UpdateCustomerServiceHistoryMultiple(data []api_structure.CustomerServiceHistoryEdit) error
	CreateCustomerServiceHistory(data api_structure.CustomerServiceHistoryForm) ([]api_structure.CustomerServiceHistoryForm, error)
	CreateCustomerServiceHistoryMultiple(data []api_structure.CustomerServiceHistoryForm) ([]api_structure.CustomerServiceHistoryForm, error)
	DeleteCustomerServiceHistory(id uuid.UUID) error
}

func (c *CustomerServiceHistoryService) GetCustomerServiceHistory(filter api_structure.CustomerServiceHistoryFilter) ([]api_structure.CustomerServiceHistory, error) {
	result := []api_structure.CustomerServiceHistory{}
	var err error
	if err = c.DB.Preload(clause.Associations).Model(&api_structure.CustomerServiceHistory{}).Where(filter).Find(&result).Error; err != nil {
		return result, err
	}
	return result, err
}

func (c *CustomerServiceHistoryService) GetCustomerServiceHistoryById(id uuid.UUID) (api_structure.CustomerServiceHistory, error) {
	result := api_structure.CustomerServiceHistory{}
	var err error
	if err = c.DB.Preload(clause.Associations).First(&result, id).Error; err != nil {
		return result, err
	}
	return result, err
}

func (c *CustomerServiceHistoryService) UpdateCustomerServiceHistory(Id uuid.UUID, data interface{}) error {
	var err error
	if err = c.DB.Model(api_structure.CustomerServiceHistory{}).Where("id = ?", Id).Updates(&data).Error; err != nil {
		return err
	}
	return err
}

func (c *CustomerServiceHistoryService) UpdateCustomerServiceHistoryMultiple(data []api_structure.CustomerServiceHistoryEdit) error {
	var err error
	for i, v := range data {
		if err = c.DB.Model(api_structure.CustomerServiceHistory{}).Where("id = ?", v.ID).Updates(&data[i]).Error; err != nil {
			return err
		}
	}
	return err
}

func (c *CustomerServiceHistoryService) CreateCustomerServiceHistory(data api_structure.CustomerServiceHistoryForm) (api_structure.CustomerServiceHistoryForm, error) {
	var err error

	if err = c.DB.Table(data.TableName()).Create(&data).Error; err != nil {
		return data, err
	}

	return data, err
}

func (c *CustomerServiceHistoryService) CreateCustomerServiceHistoryMultiple(data []api_structure.CustomerServiceHistoryForm) ([]api_structure.CustomerServiceHistoryForm, error) {
	var err error
	if err = c.DB.CreateInBatches(&data, len(data)).Error; err != nil {
		return data, err
	}
	return data, err
}

func (c *CustomerServiceHistoryService) DeleteCustomerServiceHistory(Id uuid.UUID) error {
	var err error
	if err = c.DB.Where("id = ?", Id).Delete(&api_structure.CustomerServiceHistory{}).Error; err != nil {
		return err
	}
	return err
}
