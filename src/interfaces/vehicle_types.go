package interfaces

import (
	"github.com/rochmanramadhann/fazztrack-vehicle/src/database/gorm/models"
	"github.com/rochmanramadhann/fazztrack-vehicle/src/helpers"
)

type VehicleTypeRepository interface {
	SearchVehicleTypes(vehicleType *models.VehicleType) (*models.VehicleTypes, error)
	SortVehicleTypes(sort string) (*models.VehicleTypes, error)
	GetVehicleTypes() (*models.VehicleTypes, error)
	GetVehicleType(id int) (*models.VehicleType, error)
	AddVehicleType(vehicleType *models.VehicleType) (*models.VehicleType, error)
	UpdateVehicleType(vehicleType *models.VehicleType) (*models.VehicleType, error)
	DeleteVehicleType(vehicleType *models.VehicleType) error
}

type VehicleTypeService interface {
	SearchVehicleTypes(vehicleType *models.VehicleType) (*helpers.Response, error)
	SortVehicleTypes(sort string) (*helpers.Response, error)
	GetVehicleTypes() (*helpers.Response, error)
	GetVehicleType(id int) (*helpers.Response, error)
	AddVehicleType(vehicleType *models.VehicleType) (*helpers.Response, error)
	UpdateVehicleType(vehicleType *models.VehicleType) (*helpers.Response, error)
	DeleteVehicleType(vehicleType *models.VehicleType) (*helpers.Response, error)
}
