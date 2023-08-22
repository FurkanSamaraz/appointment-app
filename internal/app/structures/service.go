package api_structures

import (
	"time"

	"github.com/google/uuid"
)

type ServiceForm struct {
	ID          uuid.UUID `gorm:"column:id;primaryKey" json:"id" swaggertype:"string"  example:"e4a3b2c1-3b2c-4a3b-a2c1-1b2c3a4b5c6d" `
	ServiceName string    `gorm:"column:service_name" json:"serviceName" binding:"required"`
	Description string    `gorm:"column:description" json:"description" binding:"required"`
	Price       float64   `gorm:"column:price" json:"price" binding:"required"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"createdAt" swaggertype:"string"  example:"1960-11-10T04:48:09Z" `
	UpdatedAt   time.Time `gorm:"column:updated_at;nullable" json:"updatedAt" validation:"omitempty" swaggertype:"string"  example:"1986-08-02T22:27:19Z" `
	DeletedAt   time.Time `gorm:"column:deleted_at;nullable" json:"deletedAt" validation:"omitempty" swaggertype:"string"  example:"1989-07-30T10:05:53Z" `
}

func (p *ServiceForm) TableName() string {
	return "public.service"
}

type Service struct {
	ID          uuid.UUID `gorm:"column:id;primaryKey" json:"id" swaggertype:"string"  example:"e4a3b2c1-3b2c-4a3b-a2c1-1b2c3a4b5c6d" `
	ServiceName string    `gorm:"column:service_name" json:"serviceName" binding:"required"`
	Description string    `gorm:"column:description" json:"description" binding:"required"`
	Price       float64   `gorm:"column:price" json:"price" binding:"required"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"createdAt" swaggertype:"string"  example:"1960-11-10T04:48:09Z" `
	UpdatedAt   time.Time `gorm:"column:updated_at;nullable" json:"updatedAt" validation:"omitempty" swaggertype:"string"  example:"1986-08-02T22:27:19Z" `
	DeletedAt   time.Time `gorm:"column:deleted_at;nullable" json:"deletedAt" validation:"omitempty" swaggertype:"string"  example:"1989-07-30T10:05:53Z" `
}

func (p *Service) TableName() string {
	return "public.service"
}

type ServiceEdit struct {
	ID          uuid.UUID `gorm:"column:id;primaryKey" json:"id" swaggertype:"string"  example:"e4a3b2c1-3b2c-4a3b-a2c1-1b2c3a4b5c6d" `
	ServiceName string
	Description string
	Price       float64
	CreatedAt   time.Time `gorm:"column:created_at" json:"createdAt" swaggertype:"string"  example:"1960-11-10T04:48:09Z" `
	UpdatedAt   time.Time `gorm:"column:updated_at;nullable" json:"updatedAt" validation:"omitempty" swaggertype:"string"  example:"1986-08-02T22:27:19Z" `
	DeletedAt   time.Time `gorm:"column:deleted_at;nullable" json:"deletedAt" validation:"omitempty" swaggertype:"string"  example:"1989-07-30T10:05:53Z" `
}

type ServiceFilter struct {
	ID          uuid.UUID `gorm:"column:id;primaryKey" json:"id" swaggertype:"string"  example:"e4a3b2c1-3b2c-4a3b-a2c1-1b2c3a4b5c6d" `
	ServiceName string
	Description string
	Price       float64
	CreatedAt   time.Time `gorm:"column:created_at" json:"createdAt" swaggertype:"string"  example:"1960-11-10T04:48:09Z" `
	UpdatedAt   time.Time `gorm:"column:updated_at;nullable" json:"updatedAt" validation:"omitempty" swaggertype:"string"  example:"1986-08-02T22:27:19Z" `
	DeletedAt   time.Time `gorm:"column:deleted_at;nullable" json:"deletedAt" validation:"omitempty" swaggertype:"string"  example:"1989-07-30T10:05:53Z" `
}

func (p *ServiceFilter) TableName() string {
	return "public.service"
}
