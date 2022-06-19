package users

import (
	"github.com/gorilla/mux"
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
	route.HandleFunc("/", controller.AddUser).Methods("POST")
	route.HandleFunc("/{id}", controller.UpdateUser).Methods("PUT")
	route.HandleFunc("/{id}", controller.DeleteUser).Methods("DELETE")
}
