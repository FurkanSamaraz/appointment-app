package api_structures

import (
	"time"

	"github.com/google/uuid"
)

type CustomerServiceHistoryForm struct {
	ID            uuid.UUID `gorm:"column:id;primaryKey" json:"id" swaggertype:"string"  example:"e4a3b2c1-3b2c-4a3b-a2c1-1b2c3a4b5c6d" `
	CustomerID    uuid.UUID `gorm:"column:customer_id;not null" json:"customerId" validation:"required" swaggertype:"string"  example:"e4a3b2c1-3b2c-4a3b-a2c1-1b2c3a4b5c6d" `
	ServiceID     uuid.UUID `gorm:"column:service_id;not null" json:"serviceId" validation:"required" swaggertype:"string"  example:"e4a3b2c1-3b2c-4a3b-a2c1-1b2c3a4b5c6d" `
	AppointmentID uuid.UUID `gorm:"column:appointment_id;not null" json:"appointmentId" validation:"required" swaggertype:"string"  example:"e4a3b2c1-3b2c-4a3b-a2c1-1b2c3a4b5c6d" `
	ServiceDate   time.Time `gorm:"column:service_date;not null" json:"serviceDate" validation:"required" swaggertype:"string"  example:"1960-11-10T04:48:09Z" `
	Notes         string    `gorm:"column:notes;not null" json:"notes" validation:"required" swaggertype:"string"  example:"Quisquam voluptatem." `
	CreatedAt     time.Time `gorm:"column:created_at" json:"createdAt" swaggertype:"string"  example:"1960-11-10T04:48:09Z" `
	UpdatedAt     time.Time `gorm:"column:updated_at;nullable" json:"updatedAt" validation:"omitempty" swaggertype:"string"  example:"1986-08-02T22:27:19Z" `
	DeletedAt     time.Time `gorm:"column:deleted_at;nullable" json:"deletedAt" validation:"omitempty" swaggertype:"string"  example:"1989-07-30T10:05:53Z" `
}

func (p *CustomerServiceHistoryForm) TableName() string {
	return "public.customer_service_history"
}

type CustomerServiceHistory struct {
	ID            uuid.UUID `gorm:"column:id;primaryKey" json:"id" swaggertype:"string"  example:"e4a3b2c1-3b2c-4a3b-a2c1-1b2c3a4b5c6d" `
	CustomerID    uuid.UUID `gorm:"column:customer_id;not null" json:"customerId" validation:"required" swaggertype:"string"  example:"e4a3b2c1-3b2c-4a3b-a2c1-1b2c3a4b5c6d" `
	ServiceID     uuid.UUID `gorm:"column:service_id;not null" json:"serviceId" validation:"required" swaggertype:"string"  example:"e4a3b2c1-3b2c-4a3b-a2c1-1b2c3a4b5c6d" `
	AppointmentID uuid.UUID `gorm:"column:appointment_id;not null" json:"appointmentId" validation:"required" swaggertype:"string"  example:"e4a3b2c1-3b2c-4a3b-a2c1-1b2c3a4b5c6d" `
	ServiceDate   time.Time `gorm:"column:service_date;not null" json:"serviceDate" validation:"required" swaggertype:"string"  example:"1960-11-10T04:48:09Z" `
	Notes         string    `gorm:"column:notes;not null" json:"notes" validation:"required" swaggertype:"string"  example:"Quisquam voluptatem." `
	CreatedAt     time.Time `gorm:"column:created_at" json:"createdAt" swaggertype:"string"  example:"1960-11-10T04:48:09Z" `
	UpdatedAt     time.Time `gorm:"column:updated_at;nullable" json:"updatedAt" validation:"omitempty" swaggertype:"string"  example:"1986-08-02T22:27:19Z" `
	DeletedAt     time.Time `gorm:"column:deleted_at;nullable" json:"deletedAt" validation:"omitempty" swaggertype:"string"  example:"1989-07-30T10:05:53Z" `
}

func (p *CustomerServiceHistory) TableName() string {
	return "public.customer_service_history"
}

type CustomerServiceHistoryEdit struct {
	ID            uuid.UUID `gorm:"column:id;primaryKey" json:"id" swaggertype:"string"  example:"e4a3b2c1-3b2c-4a3b-a2c1-1b2c3a4b5c6d" `
	CustomerID    uuid.UUID `gorm:"column:customer_id;not null" json:"customerId" validation:"required" swaggertype:"string"  example:"e4a3b2c1-3b2c-4a3b-a2c1-1b2c3a4b5c6d" `
	ServiceID     uuid.UUID `gorm:"column:service_id;not null" json:"serviceId" validation:"required" swaggertype:"string"  example:"e4a3b2c1-3b2c-4a3b-a2c1-1b2c3a4b5c6d" `
	AppointmentID uuid.UUID `gorm:"column:appointment_id;not null" json:"appointmentId" validation:"required" swaggertype:"string"  example:"e4a3b2c1-3b2c-4a3b-a2c1-1b2c3a4b5c6d" `
	ServiceDate   time.Time `gorm:"column:service_date;not null" json:"serviceDate" validation:"required" swaggertype:"string"  example:"1960-11-10T04:48:09Z" `
	Notes         string    `gorm:"column:notes;not null" json:"notes" validation:"required" swaggertype:"string"  example:"Quisquam voluptatem." `
	CreatedAt     time.Time `gorm:"column:created_at" json:"createdAt" swaggertype:"string"  example:"1960-11-10T04:48:09Z" `
	UpdatedAt     time.Time `gorm:"column:updated_at;nullable" json:"updatedAt" validation:"omitempty" swaggertype:"string"  example:"1986-08-02T22:27:19Z" `
	DeletedAt     time.Time `gorm:"column:deleted_at;nullable" json:"deletedAt" validation:"omitempty" swaggertype:"string"  example:"1989-07-30T10:05:53Z" `
}

type CustomerServiceHistoryFilter struct {
	ID            uuid.UUID `gorm:"column:id;primaryKey" json:"id" swaggertype:"string"  example:"e4a3b2c1-3b2c-4a3b-a2c1-1b2c3a4b5c6d" `
	CustomerID    uuid.UUID `gorm:"column:customer_id;not null" json:"customerId" validation:"required" swaggertype:"string"  example:"e4a3b2c1-3b2c-4a3b-a2c1-1b2c3a4b5c6d" `
	ServiceID     uuid.UUID `gorm:"column:service_id;not null" json:"serviceId" validation:"required" swaggertype:"string"  example:"e4a3b2c1-3b2c-4a3b-a2c1-1b2c3a4b5c6d" `
	AppointmentID uuid.UUID `gorm:"column:appointment_id;not null" json:"appointmentId" validation:"required" swaggertype:"string"  example:"e4a3b2c1-3b2c-4a3b-a2c1-1b2c3a4b5c6d" `
	ServiceDate   time.Time `gorm:"column:service_date;not null" json:"serviceDate" validation:"required" swaggertype:"string"  example:"1960-11-10T04:48:09Z" `
	Notes         string    `gorm:"column:notes;not null" json:"notes" validation:"required" swaggertype:"string"  example:"Quisquam voluptatem." `
	CreatedAt     time.Time `gorm:"column:created_at" json:"createdAt" swaggertype:"string"  example:"1960-11-10T04:48:09Z" `
	UpdatedAt     time.Time `gorm:"column:updated_at;nullable" json:"updatedAt" validation:"omitempty" swaggertype:"string"  example:"1986-08-02T22:27:19Z" `
	DeletedAt     time.Time `gorm:"column:deleted_at;nullable" json:"deletedAt" validation:"omitempty" swaggertype:"string"  example:"1989-07-30T10:05:53Z" `
}

func (p *CustomerServiceHistoryFilter) TableName() string {
	return "public.customer_service_history"
}
