package models

type Favorite struct {
	Id        int    `gorm:"primary_key;auto_increment" json:"id"`
	UserId    uint64 `gorm:"type:bigint;not null" json:"user_id"`
	VehicleId uint64 `gorm:"type:bigint;not null" json:"vehicle_id"`
}

type Favorites []Favorite
