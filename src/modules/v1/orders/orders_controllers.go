package orders

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rochmanramadhann/fazztrack-vehicle/src/database/gorm/models"
	"github.com/rochmanramadhann/fazztrack-vehicle/src/interfaces"
)

type OrderController struct {
	service interfaces.OrderService
}

func NewCtrl(service interfaces.OrderService) *OrderController {
	return &OrderController{service}
}

func (ctrl *OrderController) SearchOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// vars := r.URL.Query()
	var request models.Order

	// request.OrderDate = vars.Get("order_date")

	data, err := ctrl.service.SearchOrders(&request)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		data.Send(w)
	} else {
		data.Send(w)
	}
}

func (ctrl *OrderController) SortOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	sort := r.URL.Query().Get("sort")

	data, err := ctrl.service.SortOrders(sort)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		data.Send(w)
	} else {
		data.Send(w)
	}
}

func (ctrl *OrderController) GetOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := ctrl.service.GetOrders()
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		data.Send(w)
	} else {
		data.Send(w)
	}
}

func (ctrl *OrderController) GetOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	idToInt, _ := strconv.Atoi(id)

	data, err := ctrl.service.GetOrder(idToInt)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		data.Send(w)
	} else {
		data.Send(w)
	}
}

func (ctrl *OrderController) AddOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var request models.Order
	json.NewDecoder(r.Body).Decode(&request)

	fmt.Println(request)

	data, err := ctrl.service.AddOrder(&request)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		data.Send(w)
	} else {
		data.Send(w)
	}
}

func (ctrl *OrderController) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	var request models.Order
	id := vars["id"]

	json.NewDecoder(r.Body).Decode(&request)

	idToInt, _ := strconv.Atoi(id)
	request.ID = uint(idToInt)

	data, err := ctrl.service.UpdateOrder(&request)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		data.Send(w)
	} else {
		data.Send(w)
	}
}

func (ctrl *OrderController) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	var request models.Order
	id := vars["id"]

	json.NewDecoder(r.Body).Decode(&request)

	idToInt, _ := strconv.Atoi(id)
	request.ID = uint(idToInt)

	data, err := ctrl.service.DeleteOrder(&request)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		data.Send(w)
	} else {
		data.Send(w)
	}
}

// func (rep *OrderController) AddData(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	var data models.Order

// 	json.NewDecoder(r.Body).Decode(&data)
// 	result, err := rep.repo.Save(&data)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 	}

// 	result.Send(w)

// }
