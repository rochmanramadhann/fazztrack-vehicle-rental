package vehicles

import (
	"github.com/rochmanramadhann/fazztrack-vehicle/src/database/gorm/models"
	"github.com/rochmanramadhann/fazztrack-vehicle/src/helpers"
	"github.com/rochmanramadhann/fazztrack-vehicle/src/interfaces"
)

type VehicleService struct {
	repository interfaces.VehicleRepository
}

func NewVehicleService(repository interfaces.VehicleRepository) *VehicleService {
	return &VehicleService{repository}
}

func (r *VehicleService) SearchVehicles(vehicle *models.Vehicle) (*helpers.Response, error) {
	data, err := r.repository.SearchVehicles(vehicle)

	if err != nil {
		response := helpers.NewResponse(data, 400, true, err.Error())
		return response, err
	}

	if len(*data) == 0 {
		response := helpers.NewResponse(data, 404, false, "Vehicles Not Found")
		return response, nil
	}

	response := helpers.NewResponse(data, 200, false, "Vehicles Found")
	return response, nil
}

func (r *VehicleService) SortVehicles(sort string) (*helpers.Response, error) {
	data, err := r.repository.SortVehicles(sort)

	if err != nil {
		response := helpers.NewResponse(data, 400, true, err.Error())
		return response, err
	}

	if len(*data) == 0 {
		response := helpers.NewResponse(data, 404, false, "Vehicles Not Found")
		return response, nil
	}

	response := helpers.NewResponse(data, 200, false, "Vehicles Found")
	return response, nil
}

func (r *VehicleService) GetVehicles() (*helpers.Response, error) {
	data, err := r.repository.GetVehicles()

	if err != nil {
		response := helpers.NewResponse(data, 400, true, err.Error())
		return response, err
	}

	if len(*data) == 0 {
		response := helpers.NewResponse(data, 404, false, "Vehicles Not Found")
		return response, nil
	}

	response := helpers.NewResponse(data, 200, false, "Vehicles Found")
	return response, nil
}

func (r *VehicleService) GetVehicle(id int) (*helpers.Response, error) {
	data, err := r.repository.GetVehicle(id)

	if err != nil {
		response := helpers.NewResponse(data, 400, true, err.Error())
		return response, err
	}

	// if *&data.Model.ID == 0 {
	// 	response := helpers.NewResponse(data, 404, false, "Vehicle Not Found")
	// 	return response, nil
	// }

	response := helpers.NewResponse(data, 200, false, "Vehicle Found")
	return response, nil
}

func (r *VehicleService) AddVehicle(vehicle *models.Vehicle) (*helpers.Response, error) {
	data, err := r.repository.AddVehicle(vehicle)

	if err != nil {
		response := helpers.NewResponse(data, 400, true, err.Error())
		return response, err
	}

	// if *&data.Model.ID == 0 {
	// 	response := helpers.NewResponse(data, 404, false, "Vehicle Not Found")
	// 	return response, nil
	// }

	response := helpers.NewResponse(data, 200, false, "Success Add Vehicle")
	return response, nil
}

func (r *VehicleService) UpdateVehicle(vehicle *models.Vehicle) (*helpers.Response, error) {
	data, err := r.repository.UpdateVehicle(vehicle)

	if err != nil {
		response := helpers.NewResponse(data, 400, true, err.Error())
		return response, err
	}

	// if *&data.Model.ID == 0 {
	// 	response := helpers.NewResponse(data, 404, false, "Vehicle Not Found")
	// 	return response, nil
	// }

	response := helpers.NewResponse(data, 200, false, "Vehicle Found")
	return response, nil
}

func (r *VehicleService) DeleteVehicle(vehicle *models.Vehicle) (*helpers.Response, error) {
	err := r.repository.DeleteVehicle(vehicle)

	if err != nil {
		response := helpers.NewResponse(err, 400, true, err.Error())
		return response, err
	}

	// if *&data.Model.ID == 0 {
	// 	response := helpers.NewResponse(data, 404, false, "Vehicle Not Found")
	// 	return response, nil
	// }

	response := helpers.NewResponse(err, 200, false, "Success Delete Vehicle")
	return response, nil
}

// func (r *VehicleService) FindByVehiclename(Vehiclename string) (*helpers.Response, error) {
// 	data, err := r.repository.FindByVehiclename(Vehiclename)
// 	if err != nil {
// 		return nil, err
// 	}

// 	response := helpers.New(data, 200, false)
// 	return response, nil
// }
// func (r *VehicleService) Save(usr *models.Vehicle) (*helpers.Response, error) {
// 	hassPssword, err := helpers.HashPasword(usr.Password)
// 	if err != nil {
// 		return nil, err
// 	}

// 	usr.Password = hassPssword
// 	data, err := r.repository.Add(usr)
// 	if err != nil {
// 		return nil, err
// 	}

// 	response := helpers.New(data, 200, false)
// 	return response, nil
// }
