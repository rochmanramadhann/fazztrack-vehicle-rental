package vehicles

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/vehicles").Subrouter()

	repository := NewVehicleRepository(db)
	service := NewVehicleService(repository)
	controller := NewCtrl(service)

	route.HandleFunc("/search", controller.SearchVehicles).Methods("GET")
	route.HandleFunc("/sort", controller.SortVehicles).Methods("GET")
	route.HandleFunc("/", controller.GetVehicles).Methods("GET")
	route.HandleFunc("/{id}", controller.GetVehicle).Methods("GET")
	route.HandleFunc("/", controller.AddVehicle).Methods("POST")
	route.HandleFunc("/{id}", controller.UpdateVehicle).Methods("PUT")
	route.HandleFunc("/{id}", controller.DeleteVehicle).Methods("DELETE")
	// route.HandleFunc("/", middleware.Do(controller.AddVehicle, middleware.CheckAuth)).Methods("POST")
	// route.HandleFunc("/{id}", middleware.Do(controller.UpdateVehicle, middleware.CheckAuth)).Methods("PUT")
	// route.HandleFunc("/{id}", middleware.Do(controller.DeleteVehicle, middleware.CheckAuth)).Methods("DELETE")
}
