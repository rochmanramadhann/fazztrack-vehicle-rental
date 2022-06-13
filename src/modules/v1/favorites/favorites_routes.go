package favorites

import (
	"github.com/gorilla/mux"
	"github.com/rochmanramadhann/fazztrack-vehicle/src/middleware"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/favorites").Subrouter()

	repository := NewFavoriteRepository(db)
	service := NewFavoriteService(repository)
	controller := NewCtrl(service)

	route.HandleFunc("/search", controller.SearchFavorites).Methods("GET")
	route.HandleFunc("/sort", controller.SortFavorites).Methods("GET")
	route.HandleFunc("/", controller.GetFavorites).Methods("GET")
	route.HandleFunc("/{id}", controller.GetFavorite).Methods("GET")
	route.HandleFunc("/", middleware.Do(controller.AddFavorite, middleware.CheckAuth)).Methods("POST")
	route.HandleFunc("/{id}", middleware.Do(controller.UpdateFavorite, middleware.CheckAuth)).Methods("PUT")
	route.HandleFunc("/{id}", middleware.Do(controller.DeleteFavorite, middleware.CheckAuth)).Methods("DELETE")
	// route.HandleFunc("/", ctrl.AddData).Methods("POST")
}
