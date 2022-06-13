package vehicle_types

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rochmanramadhann/fazztrack-vehicle/src/database/gorm/models"
	"github.com/rochmanramadhann/fazztrack-vehicle/src/interfaces"
)

type VehicleTypeController struct {
	service interfaces.VehicleTypeService
}

func NewCtrl(service interfaces.VehicleTypeService) *VehicleTypeController {
	return &VehicleTypeController{service}
}

func (ctrl *VehicleTypeController) SearchVehicleTypes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := r.URL.Query()
	var request models.VehicleType

	request.Name = vars.Get("name")
	// request.IsHighlight = vars.Get("isHighlight")

	data, err := ctrl.service.SearchVehicleTypes(&request)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		data.Send(w)
	} else {
		data.Send(w)
	}
}

func (ctrl *VehicleTypeController) SortVehicleTypes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	sort := r.URL.Query().Get("sort")

	data, err := ctrl.service.SortVehicleTypes(sort)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		data.Send(w)
	} else {
		data.Send(w)
	}
}

func (ctrl *VehicleTypeController) GetVehicleTypes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := ctrl.service.GetVehicleTypes()
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		data.Send(w)
	} else {
		data.Send(w)
	}
}

func (ctrl *VehicleTypeController) GetVehicleType(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	idToInt, _ := strconv.Atoi(id)

	data, err := ctrl.service.GetVehicleType(idToInt)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		data.Send(w)
	} else {
		data.Send(w)
	}
}

func (ctrl *VehicleTypeController) AddVehicleType(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var request models.VehicleType
	json.NewDecoder(r.Body).Decode(&request)

	fmt.Println(request)

	data, err := ctrl.service.AddVehicleType(&request)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		data.Send(w)
	} else {
		data.Send(w)
	}
}

func (ctrl *VehicleTypeController) UpdateVehicleType(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	var request models.VehicleType
	id := vars["id"]

	json.NewDecoder(r.Body).Decode(&request)

	idToInt, _ := strconv.Atoi(id)
	request.ID = uint(idToInt)

	data, err := ctrl.service.UpdateVehicleType(&request)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		data.Send(w)
	} else {
		data.Send(w)
	}
}

func (ctrl *VehicleTypeController) DeleteVehicleType(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	var request models.VehicleType
	id := vars["id"]

	json.NewDecoder(r.Body).Decode(&request)

	idToInt, _ := strconv.Atoi(id)
	request.ID = uint(idToInt)

	data, err := ctrl.service.DeleteVehicleType(&request)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		data.Send(w)
	} else {
		data.Send(w)
	}
}

// func (rep *VehicleTypeController) AddData(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	var data models.VehicleType

// 	json.NewDecoder(r.Body).Decode(&data)
// 	result, err := rep.repo.Save(&data)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 	}

// 	result.Send(w)

// }
