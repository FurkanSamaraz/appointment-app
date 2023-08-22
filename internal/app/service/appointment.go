package service

import (
	api_structure "meeting_app/internal/app/structures"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

//go:generate mockgen -destination=../../../tests/mocks/service/mock_appointment.go -package=service meeting_app/internal/app/service IAppointment
type AppointmentService struct{ DB *gorm.DB }

type IAppointment interface {
	GetAppointment(filter api_structure.AppointmentFilter) ([]api_structure.Appointment, error)
	GetAppointmentById(id uuid.UUID) (api_structure.Appointment, error)
	UpdateAppointment(id uuid.UUID, data api_structure.AppointmentEdit) error
	UpdateAppointmentMultiple(data []api_structure.AppointmentEdit) error
	CreateAppointment(data api_structure.AppointmentForm) ([]api_structure.AppointmentForm, error)
	CreateAppointmentMultiple(data []api_structure.AppointmentForm) ([]api_structure.AppointmentForm, error)
	DeleteAppointment(id uuid.UUID) error
}

func (c *AppointmentService) GetAppointment(filter api_structure.AppointmentFilter) ([]api_structure.Appointment, error) {
	result := []api_structure.Appointment{}
	var err error
	if err = c.DB.Preload(clause.Associations).Model(&api_structure.Appointment{}).Where(filter).Find(&result).Error; err != nil {
		return result, err
	}
	return result, err
}

func (c *AppointmentService) GetAppointmentById(id uuid.UUID) (api_structure.Appointment, error) {
	result := api_structure.Appointment{}
	var err error
	if err = c.DB.Preload(clause.Associations).First(&result, id).Error; err != nil {
		return result, err
	}
	return result, err
}

func (c *AppointmentService) UpdateAppointment(Id uuid.UUID, data interface{}) error {
	var err error
	if err = c.DB.Model(api_structure.Appointment{}).Where("id = ?", Id).Updates(&data).Error; err != nil {
		return err
	}
	return err
}

func (c *AppointmentService) UpdateAppointmentMultiple(data []api_structure.AppointmentEdit) error {
	var err error
	for i, v := range data {
		if err = c.DB.Model(api_structure.Appointment{}).Where("id = ?", v.ID).Updates(&data[i]).Error; err != nil {
			return err
		}
	}
	return err
}

func (c *AppointmentService) CreateAppointment(data api_structure.AppointmentForm) (api_structure.AppointmentForm, error) {
	var err error

	if err = c.DB.Table(data.TableName()).Create(&data).Error; err != nil {
		return data, err
	}

	return data, err
}

func (c *AppointmentService) CreateAppointmentMultiple(data []api_structure.AppointmentForm) ([]api_structure.AppointmentForm, error) {
	var err error
	if err = c.DB.CreateInBatches(&data, len(data)).Error; err != nil {
		return data, err
	}
	return data, err
}

func (c *AppointmentService) DeleteAppointment(Id uuid.UUID) error {
	var err error
	if err = c.DB.Where("id = ?", Id).Delete(&api_structure.Appointment{}).Error; err != nil {
		return err
	}
	return err
}
