package models

import (
	"gorm.io/gorm"
)

type VehicleType struct {
	gorm.Model
	Name    string   `gorm:"size:100;not null" json:"name"`
	Vehicle Vehicles `gorm:"references:id" json:"vehicles,omitempty"`
}

type VehicleTypes []VehicleType
