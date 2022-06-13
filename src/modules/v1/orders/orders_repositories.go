package orders

import (
	"fmt"
	"strings"
	"time"

	"github.com/rochmanramadhann/fazztrack-vehicle/src/database/gorm/models"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db}
}

func (r *OrderRepository) SearchOrders(order *models.Order) (*models.Orders, error) {
	var orders models.Orders

	result := r.db.Where("order_date = ?", time.Now()).Find(&orders)

	if result.Error != nil {
		return nil, result.Error
	}

	return &orders, nil
}

func (r *OrderRepository) SortOrders(sort string) (*models.Orders, error) {
	var orders models.Orders

	var result *gorm.DB

	if strings.ToLower(sort) == "asc" {
		result = r.db.Order("id asc").Find(&orders)
	} else if strings.ToLower(sort) == "desc" {
		result = r.db.Order("id desc").Find(&orders)
	} else {
		result = r.db.Find(&orders)
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return &orders, nil
}

func (r *OrderRepository) GetOrders() (*models.Orders, error) {
	var orders models.Orders

	result := r.db.Find(&orders)

	if result.Error != nil {
		return nil, result.Error
	}

	return &orders, nil
}

func (r *OrderRepository) GetOrder(id int) (*models.Order, error) {
	var order models.Order

	result := r.db.First(&order, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &order, nil
}

func (r *OrderRepository) AddOrder(order *models.Order) (*models.Order, error) {
	result := r.db.Create(order)

	if result.Error != nil {
		return nil, result.Error
	}

	return order, nil
}

func (r *OrderRepository) UpdateOrder(order *models.Order) (*models.Order, error) {
	// OrderExist := r.db.First(Order, Order.Model.ID)
	// if OrderExist.Error != nil {
	// 	fmt.Println("1")
	// 	return nil, OrderExist.Error
	// }
	result := r.db.Save(order)

	if result.Error != nil {
		fmt.Println("2")
		return nil, result.Error
	}

	fmt.Println("3")
	return order, nil
}

func (r *OrderRepository) DeleteOrder(order *models.Order) error {
	orderExist := r.db.First(&order, order.Model.ID)
	if orderExist.Error != nil {
		fmt.Println("1")
		return orderExist.Error
	}
	result := r.db.Delete(&order)

	if result.Error != nil {
		fmt.Println("2")
		return result.Error
	}

	fmt.Println("3")
	return nil
}
