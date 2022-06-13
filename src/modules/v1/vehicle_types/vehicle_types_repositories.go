package vehicle_types

import (
	"fmt"
	"strings"

	"github.com/rochmanramadhann/fazztrack-vehicle/src/database/gorm/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type VehicleTypeRepository struct {
	db *gorm.DB
}

func NewVehicleTypeRepository(db *gorm.DB) *VehicleTypeRepository {
	return &VehicleTypeRepository{db}
}

func (r *VehicleTypeRepository) SearchVehicleTypes(vehicleType *models.VehicleType) (*models.VehicleTypes, error) {
	var vehicleTypes models.VehicleTypes

	result := r.db.Where("name = ?", vehicleType.Name).Preload(clause.Associations).Find(&vehicleTypes)

	if result.Error != nil {
		return nil, result.Error
	}

	return &vehicleTypes, nil
}

func (r *VehicleTypeRepository) SortVehicleTypes(sort string) (*models.VehicleTypes, error) {
	var vehicleTypes models.VehicleTypes

	var result *gorm.DB

	if strings.ToLower(sort) == "asc" {
		result = r.db.Preload(clause.Associations).Order("id asc").Find(&vehicleTypes)
	} else if strings.ToLower(sort) == "desc" {
		result = r.db.Preload(clause.Associations).Order("id desc").Find(&vehicleTypes)
	} else {
		result = r.db.Find(&vehicleTypes)
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return &vehicleTypes, nil
}

func (r *VehicleTypeRepository) GetVehicleTypes() (*models.VehicleTypes, error) {
	var vehicleTypes models.VehicleTypes

	result := r.db.Preload(clause.Associations).Find(&vehicleTypes)

	if result.Error != nil {
		return nil, result.Error
	}

	return &vehicleTypes, nil
}

func (r *VehicleTypeRepository) GetVehicleType(id int) (*models.VehicleType, error) {
	var vehicleType models.VehicleType

	result := r.db.Preload(clause.Associations).First(&vehicleType, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &vehicleType, nil
}

func (r *VehicleTypeRepository) AddVehicleType(vehicleType *models.VehicleType) (*models.VehicleType, error) {
	result := r.db.Create(vehicleType)

	if result.Error != nil {
		return nil, result.Error
	}

	return vehicleType, nil
}

func (r *VehicleTypeRepository) UpdateVehicleType(vehicleType *models.VehicleType) (*models.VehicleType, error) {
	// VehicleTypeExist := r.db.First(VehicleType, VehicleType.Model.ID)
	// if VehicleTypeExist.Error != nil {
	// 	fmt.Println("1")
	// 	return nil, VehicleTypeExist.Error
	// }
	result := r.db.Preload(clause.Associations).Save(vehicleType)

	if result.Error != nil {
		fmt.Println("2")
		return nil, result.Error
	}

	fmt.Println("3")
	return vehicleType, nil
}

func (r *VehicleTypeRepository) DeleteVehicleType(vehicleType *models.VehicleType) error {
	VehicleTypeExist := r.db.First(&vehicleType, vehicleType.Model.ID)
	if VehicleTypeExist.Error != nil {
		fmt.Println("1")
		return VehicleTypeExist.Error
	}
	result := r.db.Delete(&vehicleType)

	if result.Error != nil {
		fmt.Println("2")
		return result.Error
	}

	fmt.Println("3")
	return nil
}
