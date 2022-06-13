package vehicle_types

import (
	"github.com/gorilla/mux"
	"github.com/rochmanramadhann/fazztrack-vehicle/src/middleware"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/vehicles/types").Subrouter()

	repository := NewVehicleTypeRepository(db)
	service := NewVehicleTypeService(repository)
	controller := NewCtrl(service)

	route.HandleFunc("/search", controller.SearchVehicleTypes).Methods("GET")
	route.HandleFunc("/sort", controller.SortVehicleTypes).Methods("GET")
	route.HandleFunc("/", controller.GetVehicleTypes).Methods("GET")
	route.HandleFunc("/{id}", controller.GetVehicleType).Methods("GET")
	route.HandleFunc("/", middleware.Do(controller.AddVehicleType, middleware.CheckAuth)).Methods("POST")
	route.HandleFunc("/{id}", middleware.Do(controller.UpdateVehicleType, middleware.CheckAuth)).Methods("PUT")
	route.HandleFunc("/{id}", middleware.Do(controller.DeleteVehicleType, middleware.CheckAuth)).Methods("DELETE")
}
