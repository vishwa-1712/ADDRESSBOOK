package models

import "gorm.io/gorm"

type Contact struct {
	// gorm.Model //includes ID, CreatedAt, UpdatedAt, DeletedAt fields
	ID           uint           `gorm:"primaryKey" json:"id"`
	FirstName    string         `json:"firstname"`
	LastName     string         `json:"lastname"`
	Email        string         `json:"email" validate:"required,email"`
	Phone        *int           `json:"phone"` // pointer
	AddressLine1 string         `json:"addressline1"`
	AddressLine2 string         `json:"addressline2"`
	City         string         `json:"city"`
	State        string         `json:"state"`
	Country      string         `json:"country"`
	Pincode      *int           `json:"pincode" validate:"numeric"` // pointer
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
