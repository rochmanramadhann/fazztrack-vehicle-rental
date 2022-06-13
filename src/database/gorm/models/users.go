package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Orders      Orders    `json:"orders,omitempty"`
	Favorites   Favorites `json:"favorites,omitempty"`
	Name        string    `gorm:"size:100;not null" json:"name"`
	Email       string    `gorm:"size:100;not null" json:"email"`
	Password    string    `gorm:"size:255;not null" json:"password,omitempty"`
	PhoneNumber string    `gorm:"size:25" json:"phone_number"`
	CountryCode string    `gorm:"size:4" json:"country_code"`
	DateOfBirth string    `gorm:"size:50" json:"date_of_birth"`
	Address     string    `gorm:"size:255" json:"address"`
}

type Users []User
