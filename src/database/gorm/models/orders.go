package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserId     uint64    `gorm:"type:bigint;not null"  json:"user_id"`
	TypeId     uint64    `gorm:"type:bigint;not null" json:"type_id"`
	OrderCode  string    `gorm:"size:10;not null" json:"order_code"`
	Status     uint      `json:"status"`
	Qty        uint8     `json:"qty"`
	OrderDate  time.Time `json:"order_date"`
	ReturnDate time.Time `json:"return_date"`
}

type Orders []Order
