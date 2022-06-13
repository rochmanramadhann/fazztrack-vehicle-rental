package auth

import (
	"github.com/gorilla/mux"
	"github.com/rochmanramadhann/fazztrack-vehicle/src/modules/v1/users"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/auth").Subrouter()

	repository := users.NewUserRepository(db)
	service := NewAuthService(repository)
	controller := NewCtrl(service)

	route.HandleFunc("/", controller.Login).Methods("POST")
}
