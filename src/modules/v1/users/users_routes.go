package users

import (
	"github.com/gorilla/mux"
	"github.com/rochmanramadhann/fazztrack-vehicle/src/middleware"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/users").Subrouter()

	repository := NewUserRepository(db)
	service := NewUserService(repository)
	controller := NewCtrl(service)

	route.HandleFunc("/search", controller.SearchUsers).Methods("GET")
	route.HandleFunc("/sort", controller.SortUsers).Methods("GET")
	route.HandleFunc("/find", controller.GetUser).Methods("GET")
	route.HandleFunc("/", controller.GetUsers).Methods("GET")
	route.HandleFunc("/{id}", controller.GetUser).Methods("GET")
	route.HandleFunc("/", middleware.Do(controller.AddUser, middleware.CheckAuth)).Methods("POST")
	route.HandleFunc("/{id}", middleware.Do(controller.UpdateUser, middleware.CheckAuth)).Methods("PUT")
	route.HandleFunc("/{id}", middleware.Do(controller.DeleteUser, middleware.CheckAuth)).Methods("DELETE")
}
