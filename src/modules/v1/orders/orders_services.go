package orders

import (
	"github.com/rochmanramadhann/fazztrack-vehicle/src/database/gorm/models"
	"github.com/rochmanramadhann/fazztrack-vehicle/src/helpers"
	"github.com/rochmanramadhann/fazztrack-vehicle/src/interfaces"
)

type OrderService struct {
	repository interfaces.OrderRepository
}

func NewOrderService(repository interfaces.OrderRepository) *OrderService {
	return &OrderService{repository}
}

func (r *OrderService) SearchOrders(order *models.Order) (*helpers.Response, error) {
	data, err := r.repository.SearchOrders(order)

	if err != nil {
		response := helpers.NewResponse(data, 400, true, err.Error())
		return response, err
	}

	if len(*data) == 0 {
		response := helpers.NewResponse(data, 404, false, "Orders Not Found")
		return response, nil
	}

	response := helpers.NewResponse(data, 200, false, "Orders Found")
	return response, nil
}

func (r *OrderService) SortOrders(sort string) (*helpers.Response, error) {
	data, err := r.repository.SortOrders(sort)

	if err != nil {
		response := helpers.NewResponse(data, 400, true, err.Error())
		return response, err
	}

	if len(*data) == 0 {
		response := helpers.NewResponse(data, 404, false, "Orders Not Found")
		return response, nil
	}

	response := helpers.NewResponse(data, 200, false, "Orders Found")
	return response, nil
}

func (r *OrderService) GetOrders() (*helpers.Response, error) {
	data, err := r.repository.GetOrders()

	if err != nil {
		response := helpers.NewResponse(data, 400, true, err.Error())
		return response, err
	}

	if len(*data) == 0 {
		response := helpers.NewResponse(data, 404, false, "Orders Not Found")
		return response, nil
	}

	response := helpers.NewResponse(data, 200, false, "Orders Found")
	return response, nil
}

func (r *OrderService) GetOrder(id int) (*helpers.Response, error) {
	data, err := r.repository.GetOrder(id)

	if err != nil {
		response := helpers.NewResponse(data, 400, true, err.Error())
		return response, err
	}

	// if *&data.Model.ID == 0 {
	// 	response := helpers.NewResponse(data, 404, false, "Order Not Found")
	// 	return response, nil
	// }

	response := helpers.NewResponse(data, 200, false, "Order Found")
	return response, nil
}

func (r *OrderService) AddOrder(order *models.Order) (*helpers.Response, error) {
	data, err := r.repository.AddOrder(order)

	if err != nil {
		response := helpers.NewResponse(data, 400, true, err.Error())
		return response, err
	}

	// if *&data.Model.ID == 0 {
	// 	response := helpers.NewResponse(data, 404, false, "Order Not Found")
	// 	return response, nil
	// }

	response := helpers.NewResponse(data, 200, false, "Success Add Order")
	return response, nil
}

func (r *OrderService) UpdateOrder(order *models.Order) (*helpers.Response, error) {
	data, err := r.repository.UpdateOrder(order)

	if err != nil {
		response := helpers.NewResponse(data, 400, true, err.Error())
		return response, err
	}

	// if *&data.Model.ID == 0 {
	// 	response := helpers.NewResponse(data, 404, false, "Order Not Found")
	// 	return response, nil
	// }

	response := helpers.NewResponse(data, 200, false, "Order Found")
	return response, nil
}

func (r *OrderService) DeleteOrder(order *models.Order) (*helpers.Response, error) {
	err := r.repository.DeleteOrder(order)

	if err != nil {
		response := helpers.NewResponse(err, 400, true, err.Error())
		return response, err
	}

	// if *&data.Model.ID == 0 {
	// 	response := helpers.NewResponse(data, 404, false, "Order Not Found")
	// 	return response, nil
	// }

	response := helpers.NewResponse(err, 200, false, "Success Delete Order")
	return response, nil
}

// func (r *OrderService) FindByOrdername(Ordername string) (*helpers.Response, error) {
// 	data, err := r.repository.FindByOrdername(Ordername)
// 	if err != nil {
// 		return nil, err
// 	}

// 	response := helpers.New(data, 200, false)
// 	return response, nil
// }
// func (r *OrderService) Save(usr *models.Order) (*helpers.Response, error) {
// 	hassPssword, err := helpers.HashPasword(usr.Password)
// 	if err != nil {
// 		return nil, err
// 	}

// 	usr.Password = hassPssword
// 	data, err := r.repository.Add(usr)
// 	if err != nil {
// 		return nil, err
// 	}

// 	response := helpers.New(data, 200, false)
// 	return response, nil
// }
