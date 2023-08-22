package service

import (
	api_structure "meeting_app/internal/app/structures"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

//go:generate mockgen -destination=../../../tests/mocks/service/mockService.go -package=service meeting_app/internal/app/service IService
type ServiceService struct{ DB *gorm.DB }

type IService interface {
	GetService(filter api_structure.ServiceFilter) ([]api_structure.Service, error)
	GetServiceById(id uuid.UUID) (api_structure.Service, error)
	UpdateService(id uuid.UUID, data api_structure.ServiceEdit) error
	UpdateServiceMultiple(data []api_structure.ServiceEdit) error
	CreateService(data api_structure.ServiceForm) ([]api_structure.ServiceForm, error)
	CreateServiceMultiple(data []api_structure.ServiceForm) ([]api_structure.ServiceForm, error)
	DeleteService(id uuid.UUID) error
}

func (c *ServiceService) GetService(filter api_structure.ServiceFilter) ([]api_structure.Service, error) {
	result := []api_structure.Service{}
	var err error
	if err = c.DB.Preload(clause.Associations).Model(&api_structure.Service{}).Where(filter).Find(&result).Error; err != nil {
		return result, err
	}
	return result, err
}

func (c *ServiceService) GetServiceById(id uuid.UUID) (api_structure.Service, error) {
	result := api_structure.Service{}
	var err error
	if err = c.DB.Preload(clause.Associations).First(&result, id).Error; err != nil {
		return result, err
	}
	return result, err
}

func (c *ServiceService) UpdateService(Id uuid.UUID, data interface{}) error {
	var err error
	if err = c.DB.Model(api_structure.Service{}).Where("id = ?", Id).Updates(&data).Error; err != nil {
		return err
	}
	return err
}

func (c *ServiceService) UpdateServiceMultiple(data []api_structure.ServiceEdit) error {
	var err error
	for i, v := range data {
		if err = c.DB.Model(api_structure.Service{}).Where("id = ?", v.ID).Updates(&data[i]).Error; err != nil {
			return err
		}
	}
	return err
}

func (c *ServiceService) CreateService(data api_structure.ServiceForm) (api_structure.ServiceForm, error) {
	var err error

	if err = c.DB.Table(data.TableName()).Create(&data).Error; err != nil {
		return data, err
	}

	return data, err
}

func (c *ServiceService) CreateServiceMultiple(data []api_structure.ServiceForm) ([]api_structure.ServiceForm, error) {
	var err error
	if err = c.DB.CreateInBatches(&data, len(data)).Error; err != nil {
		return data, err
	}
	return data, err
}

func (c *ServiceService) DeleteService(Id uuid.UUID) error {
	var err error
	if err = c.DB.Where("id = ?", Id).Delete(&api_structure.Service{}).Error; err != nil {
		return err
	}
	return err
}
