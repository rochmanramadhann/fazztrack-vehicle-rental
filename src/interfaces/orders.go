package interfaces

import (
	"github.com/rochmanramadhann/fazztrack-vehicle/src/database/gorm/models"
	"github.com/rochmanramadhann/fazztrack-vehicle/src/helpers"
)

type OrderRepository interface {
	SearchOrders(order *models.Order) (*models.Orders, error)
	SortOrders(sort string) (*models.Orders, error)
	GetOrders() (*models.Orders, error)
	GetOrder(id int) (*models.Order, error)
	AddOrder(order *models.Order) (*models.Order, error)
	UpdateOrder(order *models.Order) (*models.Order, error)
	DeleteOrder(order *models.Order) error
}

type OrderService interface {
	SearchOrders(order *models.Order) (*helpers.Response, error)
	SortOrders(sort string) (*helpers.Response, error)
	GetOrders() (*helpers.Response, error)
	GetOrder(id int) (*helpers.Response, error)
	AddOrder(order *models.Order) (*helpers.Response, error)
	UpdateOrder(order *models.Order) (*helpers.Response, error)
	DeleteOrder(order *models.Order) (*helpers.Response, error)
}
