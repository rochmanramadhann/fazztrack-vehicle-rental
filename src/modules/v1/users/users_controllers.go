package users

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rochmanramadhann/fazztrack-vehicle/src/database/gorm/models"
	"github.com/rochmanramadhann/fazztrack-vehicle/src/interfaces"
)

type UserController struct {
	service interfaces.UserService
}

func NewCtrl(service interfaces.UserService) *UserController {
	return &UserController{service}
}

func (ctrl *UserController) SearchUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := r.URL.Query()
	var request models.User

	request.Name = vars.Get("name")
	request.Email = vars.Get("email")
	request.Password = vars.Get("password")

	data, err := ctrl.service.SearchUsers(&request)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		data.Send(w)
	} else {
		data.Send(w)
	}
}

func (ctrl *UserController) SortUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	sort := r.URL.Query().Get("sort")

	data, err := ctrl.service.SortUsers(sort)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		data.Send(w)
	} else {
		data.Send(w)
	}
}

func (ctrl *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := ctrl.service.GetUsers()
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		data.Send(w)
	} else {
		data.Send(w)
	}
}

func (ctrl *UserController) GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	name := r.URL.Query().Get("name")

	data, err := ctrl.service.GetUserByUsername(name)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		data.Send(w)
	} else {
		data.Send(w)
	}
}

func (ctrl *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	idToInt, _ := strconv.Atoi(id)

	data, err := ctrl.service.GetUser(idToInt)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		data.Send(w)
	} else {
		data.Send(w)
	}
}

func (ctrl *UserController) AddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var request models.User
	json.NewDecoder(r.Body).Decode(&request)

	fmt.Println(request)

	data, err := ctrl.service.AddUser(&request)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		data.Send(w)
	} else {
		data.Send(w)
	}
}

func (ctrl *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	var request models.User
	id := vars["id"]

	json.NewDecoder(r.Body).Decode(&request)

	idToInt, _ := strconv.Atoi(id)
	request.ID = uint(idToInt)

	data, err := ctrl.service.UpdateUser(&request)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		data.Send(w)
	} else {
		data.Send(w)
	}
}

func (ctrl *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	var request models.User
	id := vars["id"]

	json.NewDecoder(r.Body).Decode(&request)

	idToInt, _ := strconv.Atoi(id)
	request.ID = uint(idToInt)

	data, err := ctrl.service.DeleteUser(&request)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		data.Send(w)
	} else {
		data.Send(w)
	}
}
