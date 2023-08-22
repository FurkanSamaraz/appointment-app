package api_structures

import (
	"time"

	"github.com/google/uuid"
)

type AppointmentForm struct {
	ID              uuid.UUID `gorm:"column:id;primaryKey" json:"id" swaggertype:"string"  example:"e4a3b2c1-3b2c-4a3b-a2c1-1b2c3a4b5c6d" `
	CustomerID      uuid.UUID `gorm:"column:customer_id" json:"customerID" swaggertype:"string"  example:"e4a3b2c1-3b2c-4a3b-a2c1-1b2c3a4b5c6d" `
	AppointmentDate time.Time `gorm:"column:appointment_date" json:"appointmentDate" swaggertype:"string"  example:"1960-11-10T04:48:09Z" `
	Purpose         string    `gorm:"column:purpose" json:"purpose" swaggertype:"string"  example:"1960-11-10T04:48:09Z" `
	CreatedAt       time.Time `gorm:"column:created_at" json:"createdAt" swaggertype:"string"  example:"1960-11-10T04:48:09Z" `
	UpdatedAt       time.Time `gorm:"column:updated_at;nullable" json:"updatedAt" validation:"omitempty" swaggertype:"string"  example:"1986-08-02T22:27:19Z" `
	DeletedAt       time.Time `gorm:"column:deleted_at;nullable" json:"deletedAt" validation:"omitempty" swaggertype:"string"  example:"1989-07-30T10:05:53Z" `
}

func (p *AppointmentForm) TableName() string {
	return "public.appointment"
}

type Appointment struct {
	ID              uuid.UUID `gorm:"column:id;primaryKey" json:"id" swaggertype:"string"  example:"e4a3b2c1-3b2c-4a3b-a2c1-1b2c3a4b5c6d" `
	CustomerID      uuid.UUID `gorm:"column:customer_id" json:"customerID" swaggertype:"string"  example:"e4a3b2c1-3b2c-4a3b-a2c1-1b2c3a4b5c6d" `
	AppointmentDate time.Time `gorm:"column:appointment_date" json:"appointmentDate" swaggertype:"string"  example:"1960-11-10T04:48:09Z" `
	Purpose         string    `gorm:"column:purpose" json:"purpose" swaggertype:"string"  example:"1960-11-10T04:48:09Z" `
	CreatedAt       time.Time `gorm:"column:created_at" json:"createdAt" swaggertype:"string"  example:"1960-11-10T04:48:09Z" `
	UpdatedAt       time.Time `gorm:"column:updated_at;nullable" json:"updatedAt" validation:"omitempty" swaggertype:"string"  example:"1986-08-02T22:27:19Z" `
	DeletedAt       time.Time `gorm:"column:deleted_at;nullable" json:"deletedAt" validation:"omitempty" swaggertype:"string"  example:"1989-07-30T10:05:53Z" `
}

func (p *Appointment) TableName() string {
	return "public.appointment"
}

type AppointmentEdit struct {
	ID              uuid.UUID `gorm:"column:id;primaryKey" json:"id" swaggertype:"string"  example:"e4a3b2c1-3b2c-4a3b-a2c1-1b2c3a4b5c6d" `
	CustomerID      uuid.UUID `gorm:"column:customer_id" json:"customerID" swaggertype:"string"  example:"e4a3b2c1-3b2c-4a3b-a2c1-1b2c3a4b5c6d" `
	AppointmentDate time.Time `gorm:"column:appointment_date" json:"appointmentDate" swaggertype:"string"  example:"1960-11-10T04:48:09Z" `
	Purpose         string    `gorm:"column:purpose" json:"purpose" swaggertype:"string"  example:"1960-11-10T04:48:09Z" `
	CreatedAt       time.Time `gorm:"column:created_at" json:"createdAt" swaggertype:"string"  example:"1960-11-10T04:48:09Z" `
	UpdatedAt       time.Time `gorm:"column:updated_at;nullable" json:"updatedAt" validation:"omitempty" swaggertype:"string"  example:"1986-08-02T22:27:19Z" `
	DeletedAt       time.Time `gorm:"column:deleted_at;nullable" json:"deletedAt" validation:"omitempty" swaggertype:"string"  example:"1989-07-30T10:05:53Z" `
}

type AppointmentFilter struct {
	ID              uuid.UUID `gorm:"column:id;primaryKey" json:"id" swaggertype:"string"  example:"e4a3b2c1-3b2c-4a3b-a2c1-1b2c3a4b5c6d" `
	CustomerID      uuid.UUID `gorm:"column:customer_id" json:"customerID" swaggertype:"string"  example:"e4a3b2c1-3b2c-4a3b-a2c1-1b2c3a4b5c6d" `
	AppointmentDate time.Time `gorm:"column:appointment_date" json:"appointmentDate" swaggertype:"string"  example:"1960-11-10T04:48:09Z" `
	Purpose         string    `gorm:"column:purpose" json:"purpose" swaggertype:"string"  example:"1960-11-10T04:48:09Z" `
	CreatedAt       time.Time `gorm:"column:created_at" json:"createdAt" swaggertype:"string"  example:"1960-11-10T04:48:09Z" `
	UpdatedAt       time.Time `gorm:"column:updated_at;nullable" json:"updatedAt" validation:"omitempty" swaggertype:"string"  example:"1986-08-02T22:27:19Z" `
	DeletedAt       time.Time `gorm:"column:deleted_at;nullable" json:"deletedAt" validation:"omitempty" swaggertype:"string"  example:"1989-07-30T10:05:53Z" `
}

func (p *AppointmentFilter) TableName() string {
	return "public.appointment"
}
