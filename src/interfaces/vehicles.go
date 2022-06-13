package interfaces

import (
	"github.com/rochmanramadhann/fazztrack-vehicle/src/database/gorm/models"
	"github.com/rochmanramadhann/fazztrack-vehicle/src/helpers"
)

type VehicleRepository interface {
	SearchVehicles(vehicle *models.Vehicle) (*models.Vehicles, error)
	SortVehicles(sort string) (*models.Vehicles, error)
	GetVehicles() (*models.Vehicles, error)
	GetVehicle(id int) (*models.Vehicle, error)
	AddVehicle(vehicle *models.Vehicle) (*models.Vehicle, error)
	UpdateVehicle(vehicle *models.Vehicle) (*models.Vehicle, error)
	DeleteVehicle(vehicle *models.Vehicle) error
}

type VehicleService interface {
	SearchVehicles(vehicle *models.Vehicle) (*helpers.Response, error)
	SortVehicles(sort string) (*helpers.Response, error)
	GetVehicles() (*helpers.Response, error)
	GetVehicle(id int) (*helpers.Response, error)
	AddVehicle(vehicle *models.Vehicle) (*helpers.Response, error)
	UpdateVehicle(vehicle *models.Vehicle) (*helpers.Response, error)
	DeleteVehicle(vehicle *models.Vehicle) (*helpers.Response, error)
}
