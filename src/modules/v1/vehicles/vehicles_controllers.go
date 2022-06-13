package vehicles

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rochmanramadhann/fazztrack-vehicle/src/database/gorm/models"
	"github.com/rochmanramadhann/fazztrack-vehicle/src/interfaces"
)

type VehicleController struct {
	service interfaces.VehicleService
}

func NewCtrl(service interfaces.VehicleService) *VehicleController {
	return &VehicleController{service}
}

func (ctrl *VehicleController) SearchVehicles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := r.URL.Query()
	var request models.Vehicle

	request.Name = vars.Get("name")
	reqIsHightlight := vars.Get("is_highlight")

	request.IsHighlight, _ = strconv.ParseUint(reqIsHightlight, 10, 64)

	data, err := ctrl.service.SearchVehicles(&request)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		data.Send(w)
	} else {
		data.Send(w)
	}
}

func (ctrl *VehicleController) SortVehicles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	sort := r.URL.Query().Get("sort")

	data, err := ctrl.service.SortVehicles(sort)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		data.Send(w)
	} else {
		data.Send(w)
	}
}

func (ctrl *VehicleController) GetVehicles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := ctrl.service.GetVehicles()
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		data.Send(w)
	} else {
		data.Send(w)
	}
}

func (ctrl *VehicleController) GetVehicle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	idToInt, _ := strconv.Atoi(id)

	data, err := ctrl.service.GetVehicle(idToInt)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		data.Send(w)
	} else {
		data.Send(w)
	}
}

func (ctrl *VehicleController) AddVehicle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var request models.Vehicle
	json.NewDecoder(r.Body).Decode(&request)

	fmt.Println(request)

	data, err := ctrl.service.AddVehicle(&request)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		data.Send(w)
	} else {
		data.Send(w)
	}
}

func (ctrl *VehicleController) UpdateVehicle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	var request models.Vehicle
	id := vars["id"]

	json.NewDecoder(r.Body).Decode(&request)

	idToInt, _ := strconv.Atoi(id)
	request.ID = uint(idToInt)

	data, err := ctrl.service.UpdateVehicle(&request)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		data.Send(w)
	} else {
		data.Send(w)
	}
}

func (ctrl *VehicleController) DeleteVehicle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	var request models.Vehicle
	id := vars["id"]

	json.NewDecoder(r.Body).Decode(&request)

	idToInt, _ := strconv.Atoi(id)
	request.ID = uint(idToInt)

	data, err := ctrl.service.DeleteVehicle(&request)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		data.Send(w)
	} else {
		data.Send(w)
	}
}

// func (rep *VehicleController) AddData(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	var data models.Vehicle

// 	json.NewDecoder(r.Body).Decode(&data)
// 	result, err := rep.repo.Save(&data)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 	}

// 	result.Send(w)

// }
