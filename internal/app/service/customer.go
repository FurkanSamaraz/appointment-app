package service

import (
	api_structure "meeting_app/internal/app/structures"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

//go:generate mockgen -destination=../../../tests/mocks/service/mockCustomer.go -package=service meeting_app/internal/app/service ICustomer
type CustomerService struct{ DB *gorm.DB }

type ICustomer interface {
	GetCustomer(filter api_structure.CustomerFilter) ([]api_structure.Customer, error)
	GetCustomerById(id uuid.UUID) (api_structure.Customer, error)
	UpdateCustomer(id uuid.UUID, data api_structure.CustomerEdit) error
	UpdateCustomerMultiple(data []api_structure.CustomerEdit) error
	CreateCustomer(data api_structure.CustomerForm) ([]api_structure.CustomerForm, error)
	CreateCustomerMultiple(data []api_structure.CustomerForm) ([]api_structure.CustomerForm, error)
	DeleteCustomer(id uuid.UUID) error
}

func (c *CustomerService) GetCustomer(filter api_structure.CustomerFilter) ([]api_structure.Customer, error) {
	result := []api_structure.Customer{}
	var err error
	if err = c.DB.Preload(clause.Associations).Model(&api_structure.Customer{}).Where(filter).Find(&result).Error; err != nil {
		return result, err
	}
	return result, err
}

func (c *CustomerService) GetCustomerById(id uuid.UUID) (api_structure.Customer, error) {
	result := api_structure.Customer{}
	var err error
	if err = c.DB.Preload(clause.Associations).First(&result, id).Error; err != nil {
		return result, err
	}
	return result, err
}

func (c *CustomerService) UpdateCustomer(Id uuid.UUID, data interface{}) error {
	var err error
	if err = c.DB.Model(api_structure.Customer{}).Where("id = ?", Id).Updates(&data).Error; err != nil {
		return err
	}
	return err
}

func (c *CustomerService) UpdateCustomerMultiple(data []api_structure.CustomerEdit) error {
	var err error
	for i, v := range data {
		if err = c.DB.Model(api_structure.Customer{}).Where("id = ?", v.ID).Updates(&data[i]).Error; err != nil {
			return err
		}
	}
	return err
}

func (c *CustomerService) CreateCustomer(data api_structure.CustomerForm) (api_structure.CustomerForm, error) {
	var err error

	if err = c.DB.Table(data.TableName()).Create(&data).Error; err != nil {
		return data, err
	}

	return data, err
}

func (c *CustomerService) CreateCustomerMultiple(data []api_structure.CustomerForm) ([]api_structure.CustomerForm, error) {
	var err error
	if err = c.DB.CreateInBatches(&data, len(data)).Error; err != nil {
		return data, err
	}
	return data, err
}

func (c *CustomerService) DeleteCustomer(Id uuid.UUID) error {
	var err error
	if err = c.DB.Where("id = ?", Id).Delete(&api_structure.Customer{}).Error; err != nil {
		return err
	}
	return err
}
