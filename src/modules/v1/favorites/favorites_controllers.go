package favorites

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rochmanramadhann/fazztrack-vehicle/src/database/gorm/models"
	"github.com/rochmanramadhann/fazztrack-vehicle/src/interfaces"
)

type FavoriteController struct {
	service interfaces.FavoriteService
}

func NewCtrl(service interfaces.FavoriteService) *FavoriteController {
	return &FavoriteController{service}
}

func (ctrl *FavoriteController) SearchFavorites(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// vars := r.URL.Query()
	var request models.Favorite

	// request.FavoriteDate = vars.Get("Favorite_date")

	data, err := ctrl.service.SearchFavorites(&request)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		data.Send(w)
	} else {
		data.Send(w)
	}
}

func (ctrl *FavoriteController) SortFavorites(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	sort := r.URL.Query().Get("sort")

	data, err := ctrl.service.SortFavorites(sort)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		data.Send(w)
	} else {
		data.Send(w)
	}
}

func (ctrl *FavoriteController) GetFavorites(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := ctrl.service.GetFavorites()
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		data.Send(w)
	} else {
		data.Send(w)
	}
}

func (ctrl *FavoriteController) GetFavorite(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	idToInt, _ := strconv.Atoi(id)

	data, err := ctrl.service.GetFavorite(idToInt)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		data.Send(w)
	} else {
		data.Send(w)
	}
}

func (ctrl *FavoriteController) AddFavorite(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var request models.Favorite
	json.NewDecoder(r.Body).Decode(&request)

	fmt.Println(request)

	data, err := ctrl.service.AddFavorite(&request)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		data.Send(w)
	} else {
		data.Send(w)
	}
}

func (ctrl *FavoriteController) UpdateFavorite(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	var request models.Favorite
	id := vars["id"]

	json.NewDecoder(r.Body).Decode(&request)

	idToInt, _ := strconv.Atoi(id)
	request.Id = int(idToInt)

	data, err := ctrl.service.UpdateFavorite(&request)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		data.Send(w)
	} else {
		data.Send(w)
	}
}

func (ctrl *FavoriteController) DeleteFavorite(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	var request models.Favorite
	id := vars["id"]

	json.NewDecoder(r.Body).Decode(&request)

	idToInt, _ := strconv.Atoi(id)
	request.Id = int(idToInt)

	data, err := ctrl.service.DeleteFavorite(&request)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		data.Send(w)
	} else {
		data.Send(w)
	}
}

// func (rep *FavoriteController) AddData(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	var data models.Favorite

// 	json.NewDecoder(r.Body).Decode(&data)
// 	result, err := rep.repo.Save(&data)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 	}

// 	result.Send(w)

// }
