package orders

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/orders").Subrouter()

	repository := NewOrderRepository(db)
	service := NewOrderService(repository)
	controller := NewCtrl(service)

	route.HandleFunc("/search", controller.SearchOrders).Methods("GET")
	route.HandleFunc("/sort", controller.SortOrders).Methods("GET")
	route.HandleFunc("/", controller.GetOrders).Methods("GET")
	route.HandleFunc("/{id}", controller.GetOrder).Methods("GET")
	route.HandleFunc("/", controller.AddOrder).Methods("POST")
	route.HandleFunc("/{id}", controller.UpdateOrder).Methods("PUT")
	route.HandleFunc("/{id}", controller.DeleteOrder).Methods("DELETE")
	// route.HandleFunc("/", middleware.Do(controller.AddOrder, middleware.CheckAuth)).Methods("POST")
	// route.HandleFunc("/{id}", middleware.Do(controller.UpdateOrder, middleware.CheckAuth)).Methods("PUT")
	// route.HandleFunc("/{id}", middleware.Do(controller.DeleteOrder, middleware.CheckAuth)).Methods("DELETE")
}
