package models

import (
	"gorm.io/gorm"
)

type Vehicle struct {
	gorm.Model
	VehicleTypeId    uint64 `gorm:"type:bigint;not null" json:"vehicle_type_id"`
	Name             string `gorm:"size:100;not null" json:"name"`
	Image            string `gorm:"size:255" json:"image"`
	BookCount        uint64 `gorm:"type:bigint;not null" json:"book_count"`
	IsHighlight      uint64 `gorm:"type:bigint;not null" json:"is_highlight"`
	Stock            uint64 `gorm:"type:bigint;not null" json:"stock"`
	Price            uint64 `gorm:"type:bigint;not null" json:"price"`
	AlternativePrice uint64 `gorm:"type:bigint;not null" json:"alternative_price"`
	Longitude        string `gorm:"size:60" json:"longitude"`
	Latitude         string `gorm:"size:60" json:"latitude"`
}

type Vehicles []Vehicle
