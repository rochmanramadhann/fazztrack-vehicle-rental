package vehicle_types

import (
	"github.com/rochmanramadhann/fazztrack-vehicle/src/database/gorm/models"
	"github.com/rochmanramadhann/fazztrack-vehicle/src/helpers"
	"github.com/rochmanramadhann/fazztrack-vehicle/src/interfaces"
)

type VehicleTypeService struct {
	repository interfaces.VehicleTypeRepository
}

func NewVehicleTypeService(repository interfaces.VehicleTypeRepository) *VehicleTypeService {
	return &VehicleTypeService{repository}
}

func (r *VehicleTypeService) SearchVehicleTypes(vehicleType *models.VehicleType) (*helpers.Response, error) {
	data, err := r.repository.SearchVehicleTypes(vehicleType)

	if err != nil {
		response := helpers.NewResponse(data, 400, true, err.Error())
		return response, err
	}

	if len(*data) == 0 {
		response := helpers.NewResponse(data, 404, false, "VehicleTypes Not Found")
		return response, nil
	}

	response := helpers.NewResponse(data, 200, false, "VehicleTypes Found")
	return response, nil
}

func (r *VehicleTypeService) SortVehicleTypes(sort string) (*helpers.Response, error) {
	data, err := r.repository.SortVehicleTypes(sort)

	if err != nil {
		response := helpers.NewResponse(data, 400, true, err.Error())
		return response, err
	}

	if len(*data) == 0 {
		response := helpers.NewResponse(data, 404, false, "VehicleTypes Not Found")
		return response, nil
	}

	response := helpers.NewResponse(data, 200, false, "VehicleTypes Found")
	return response, nil
}

func (r *VehicleTypeService) GetVehicleTypes() (*helpers.Response, error) {
	data, err := r.repository.GetVehicleTypes()

	if err != nil {
		response := helpers.NewResponse(data, 400, true, err.Error())
		return response, err
	}

	if len(*data) == 0 {
		response := helpers.NewResponse(data, 404, false, "VehicleTypes Not Found")
		return response, nil
	}

	response := helpers.NewResponse(data, 200, false, "VehicleTypes Found")
	return response, nil
}

func (r *VehicleTypeService) GetVehicleType(id int) (*helpers.Response, error) {
	data, err := r.repository.GetVehicleType(id)

	if err != nil {
		response := helpers.NewResponse(data, 400, true, err.Error())
		return response, err
	}

	// if *&data.Model.ID == 0 {
	// 	response := helpers.NewResponse(data, 404, false, "VehicleType Not Found")
	// 	return response, nil
	// }

	response := helpers.NewResponse(data, 200, false, "VehicleType Found")
	return response, nil
}

func (r *VehicleTypeService) AddVehicleType(vehicleType *models.VehicleType) (*helpers.Response, error) {
	data, err := r.repository.AddVehicleType(vehicleType)

	if err != nil {
		response := helpers.NewResponse(data, 400, true, err.Error())
		return response, err
	}

	// if *&data.Model.ID == 0 {
	// 	response := helpers.NewResponse(data, 404, false, "VehicleType Not Found")
	// 	return response, nil
	// }

	response := helpers.NewResponse(data, 200, false, "Success Add VehicleType")
	return response, nil
}

func (r *VehicleTypeService) UpdateVehicleType(vehicleType *models.VehicleType) (*helpers.Response, error) {
	data, err := r.repository.UpdateVehicleType(vehicleType)

	if err != nil {
		response := helpers.NewResponse(data, 400, true, err.Error())
		return response, err
	}

	// if *&data.Model.ID == 0 {
	// 	response := helpers.NewResponse(data, 404, false, "VehicleType Not Found")
	// 	return response, nil
	// }

	response := helpers.NewResponse(data, 200, false, "VehicleType Found")
	return response, nil
}

func (r *VehicleTypeService) DeleteVehicleType(vehicleType *models.VehicleType) (*helpers.Response, error) {
	err := r.repository.DeleteVehicleType(vehicleType)

	if err != nil {
		response := helpers.NewResponse(err, 400, true, err.Error())
		return response, err
	}

	// if *&data.Model.ID == 0 {
	// 	response := helpers.NewResponse(data, 404, false, "VehicleType Not Found")
	// 	return response, nil
	// }

	response := helpers.NewResponse(err, 200, false, "Success Delete VehicleType")
	return response, nil
}

// func (r *VehicleTypeService) FindByVehicleTypename(VehicleTypename string) (*helpers.Response, error) {
// 	data, err := r.repository.FindByVehicleTypename(VehicleTypename)
// 	if err != nil {
// 		return nil, err
// 	}

// 	response := helpers.New(data, 200, false)
// 	return response, nil
// }
// func (r *VehicleTypeService) Save(usr *models.VehicleType) (*helpers.Response, error) {
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
