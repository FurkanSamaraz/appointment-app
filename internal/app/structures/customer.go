package api_structures

import (
	"time"

	"github.com/google/uuid"
)

type CustomerForm struct {
	ID          uuid.UUID `gorm:"column:id;primaryKey" json:"id" swaggertype:"string"  example:"e4a3b2c1-3b2c-4a3b-a2c1-1b2c3a4b5c6d" `
	FirstName   string    `gorm:"column:first_name;not null" json:"firstName" validation:"required" swaggertype:"string"  example:"Neque magni." `
	LastName    string    `gorm:"column:last_name;not null" json:"lastName" validation:"required" swaggertype:"string"  example:"Dignissimos rerum." `
	Email       string    `gorm:"column:email;not null" json:"email" validation:"required" swaggertype:"string"  example:"Saepe excepturi." `
	PhoneNumber string    `gorm:"column:phone_number;not null" json:"phoneNumber" validation:"required" swaggertype:"string"  example:"Sunt autem." `
	CreatedAt   time.Time `gorm:"column:created_at" json:"createdAt" swaggertype:"string"  example:"1960-11-10T04:48:09Z" `
	UpdatedAt   time.Time `gorm:"column:updated_at;nullable" json:"updatedAt" validation:"omitempty" swaggertype:"string"  example:"1986-08-02T22:27:19Z" `
	DeletedAt   time.Time `gorm:"column:deleted_at;nullable" json:"deletedAt" validation:"omitempty" swaggertype:"string"  example:"1989-07-30T10:05:53Z" `
}

func (p *CustomerForm) TableName() string {
	return "public.customer"
}

type Customer struct {
	ID          uuid.UUID `gorm:"column:id;primaryKey" json:"id" swaggertype:"string"  example:"e4a3b2c1-3b2c-4a3b-a2c1-1b2c3a4b5c6d" `
	FirstName   string    `gorm:"column:first_name;not null" json:"firstName" validation:"required" swaggertype:"string"  example:"Neque magni." `
	LastName    string    `gorm:"column:last_name;not null" json:"lastName" validation:"required" swaggertype:"string"  example:"Dignissimos rerum." `
	Email       string    `gorm:"column:email;not null" json:"email" validation:"required" swaggertype:"string"  example:"Saepe excepturi." `
	PhoneNumber string    `gorm:"column:phone_number;not null" json:"phoneNumber" validation:"required" swaggertype:"string"  example:"Sunt autem." `
	CreatedAt   time.Time `gorm:"column:created_at" json:"createdAt" swaggertype:"string"  example:"1960-11-10T04:48:09Z" `
	UpdatedAt   time.Time `gorm:"column:updated_at;nullable" json:"updatedAt" validation:"omitempty" swaggertype:"string"  example:"1986-08-02T22:27:19Z" `
	DeletedAt   time.Time `gorm:"column:deleted_at;nullable" json:"deletedAt" validation:"omitempty" swaggertype:"string"  example:"1989-07-30T10:05:53Z" `
}

func (p *Customer) TableName() string {
	return "public.customer"
}

type CustomerEdit struct {
	ID          uuid.UUID `gorm:"column:id;primaryKey" json:"id" swaggertype:"string"  example:"e4a3b2c1-3b2c-4a3b-a2c1-1b2c3a4b5c6d" `
	FirstName   string    `gorm:"column:first_name;not null" json:"firstName" validation:"required" swaggertype:"string"  example:"Neque magni." `
	LastName    string    `gorm:"column:last_name;not null" json:"lastName" validation:"required" swaggertype:"string"  example:"Dignissimos rerum." `
	Email       string    `gorm:"column:email;not null" json:"email" validation:"required" swaggertype:"string"  example:"Saepe excepturi." `
	PhoneNumber string    `gorm:"column:phone_number;not null" json:"phoneNumber" validation:"required" swaggertype:"string"  example:"Sunt autem." `
	CreatedAt   time.Time `gorm:"column:created_at" json:"createdAt" swaggertype:"string"  example:"1960-11-10T04:48:09Z" `
	UpdatedAt   time.Time `gorm:"column:updated_at;nullable" json:"updatedAt" validation:"omitempty" swaggertype:"string"  example:"1986-08-02T22:27:19Z" `
	DeletedAt   time.Time `gorm:"column:deleted_at;nullable" json:"deletedAt" validation:"omitempty" swaggertype:"string"  example:"1989-07-30T10:05:53Z" `
}

type CustomerFilter struct {
	ID          uuid.UUID `gorm:"column:id;primaryKey" json:"id" swaggertype:"string"  example:"e4a3b2c1-3b2c-4a3b-a2c1-1b2c3a4b5c6d" `
	FirstName   string    `gorm:"column:first_name;not null" json:"firstName" validation:"required" swaggertype:"string"  example:"Neque magni." `
	LastName    string    `gorm:"column:last_name;not null" json:"lastName" validation:"required" swaggertype:"string"  example:"Dignissimos rerum." `
	Email       string    `gorm:"column:email;not null" json:"email" validation:"required" swaggertype:"string"  example:"Saepe excepturi." `
	PhoneNumber string    `gorm:"column:phone_number;not null" json:"phoneNumber" validation:"required" swaggertype:"string"  example:"Sunt autem." `
	CreatedAt   time.Time `gorm:"column:created_at" json:"createdAt" swaggertype:"string"  example:"1960-11-10T04:48:09Z" `
	UpdatedAt   time.Time `gorm:"column:updated_at;nullable" json:"updatedAt" validation:"omitempty" swaggertype:"string"  example:"1986-08-02T22:27:19Z" `
	DeletedAt   time.Time `gorm:"column:deleted_at;nullable" json:"deletedAt" validation:"omitempty" swaggertype:"string"  example:"1989-07-30T10:05:53Z" `
}

func (p *CustomerFilter) TableName() string {
	return "public.customer"

}
