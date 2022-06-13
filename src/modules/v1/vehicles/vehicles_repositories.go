package vehicles

import (
	"fmt"
	"strings"

	"github.com/rochmanramadhann/fazztrack-vehicle/src/database/gorm/models"
	"gorm.io/gorm"
)

type VehicleRepository struct {
	db *gorm.DB
}

func NewVehicleRepository(db *gorm.DB) *VehicleRepository {
	return &VehicleRepository{db}
}

func (r *VehicleRepository) SearchVehicles(vehicle *models.Vehicle) (*models.Vehicles, error) {
	var vehicles models.Vehicles

	result := r.db.Where("name = ?", vehicle.Name).Or("is_highlight = ?", vehicle.IsHighlight).Find(&vehicles)

	if result.Error != nil {
		return nil, result.Error
	}

	return &vehicles, nil
}

func (r *VehicleRepository) SortVehicles(sort string) (*models.Vehicles, error) {
	var vehicles models.Vehicles

	var result *gorm.DB

	if strings.ToLower(sort) == "asc" {
		result = r.db.Order("id asc").Find(&vehicles)
	} else if strings.ToLower(sort) == "desc" {
		result = r.db.Order("id desc").Find(&vehicles)
	} else {
		result = r.db.Find(&vehicles)
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return &vehicles, nil
}

func (r *VehicleRepository) GetVehicles() (*models.Vehicles, error) {
	var vehicles models.Vehicles

	result := r.db.Find(&vehicles)

	if result.Error != nil {
		return nil, result.Error
	}

	return &vehicles, nil
}

func (r *VehicleRepository) GetVehicle(id int) (*models.Vehicle, error) {
	var vehicle models.Vehicle

	result := r.db.First(&vehicle, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &vehicle, nil
}

func (r *VehicleRepository) AddVehicle(vehicle *models.Vehicle) (*models.Vehicle, error) {
	result := r.db.Create(vehicle)

	if result.Error != nil {
		return nil, result.Error
	}

	return vehicle, nil
}

func (r *VehicleRepository) UpdateVehicle(vehicle *models.Vehicle) (*models.Vehicle, error) {
	// VehicleExist := r.db.First(Vehicle, Vehicle.Model.ID)
	// if VehicleExist.Error != nil {
	// 	fmt.Println("1")
	// 	return nil, VehicleExist.Error
	// }
	result := r.db.Save(vehicle)

	if result.Error != nil {
		fmt.Println("2")
		return nil, result.Error
	}

	fmt.Println("3")
	return vehicle, nil
}

func (r *VehicleRepository) DeleteVehicle(vehicle *models.Vehicle) error {
	VehicleExist := r.db.First(&vehicle, vehicle.Model.ID)
	if VehicleExist.Error != nil {
		fmt.Println("1")
		return VehicleExist.Error
	}
	result := r.db.Delete(&vehicle)

	if result.Error != nil {
		fmt.Println("2")
		return result.Error
	}

	fmt.Println("3")
	return nil
}
