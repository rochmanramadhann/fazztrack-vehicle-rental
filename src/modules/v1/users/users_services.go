package users

import (
	"github.com/rochmanramadhann/fazztrack-vehicle/src/database/gorm/models"
	"github.com/rochmanramadhann/fazztrack-vehicle/src/helpers"
	"github.com/rochmanramadhann/fazztrack-vehicle/src/interfaces"
)

type UserService struct {
	repository interfaces.UserRepository
}

func NewUserService(repository interfaces.UserRepository) *UserService {
	return &UserService{repository}
}

func (r *UserService) SearchUsers(user *models.User) (*helpers.Response, error) {
	data, err := r.repository.SearchUsers(user)

	if err != nil {
		response := helpers.NewResponse(data, 400, true, err.Error())
		return response, err
	}

	if len(*data) == 0 {
		response := helpers.NewResponse(data, 404, false, "Users Not Found")
		return response, nil
	}

	response := helpers.NewResponse(data, 200, false, "Users Found")
	return response, nil
}

func (r *UserService) SortUsers(sort string) (*helpers.Response, error) {
	data, err := r.repository.SortUsers(sort)

	if err != nil {
		response := helpers.NewResponse(data, 400, true, err.Error())
		return response, err
	}

	if len(*data) == 0 {
		response := helpers.NewResponse(data, 404, false, "Users Not Found")
		return response, nil
	}

	response := helpers.NewResponse(data, 200, false, "Users Found")
	return response, nil
}

func (r *UserService) GetUsers() (*helpers.Response, error) {
	data, err := r.repository.GetUsers()

	if err != nil {
		response := helpers.NewResponse(data, 400, true, err.Error())
		return response, err
	}

	if len(*data) == 0 {
		response := helpers.NewResponse(data, 404, false, "Users Not Found")
		return response, nil
	}

	response := helpers.NewResponse(data, 200, false, "Users Found")
	return response, nil
}

func (r *UserService) GetUserByUsername(username string) (*helpers.Response, error) {
	data, err := r.repository.GetUserByUsername(username)
	if err != nil {
		response := helpers.NewResponse(err, 400, true, err.Error())
		return response, err
	}

	response := helpers.NewResponse(data, 200, false, "Success Delete User")
	return response, nil
}

func (r *UserService) GetUser(id int) (*helpers.Response, error) {
	data, err := r.repository.GetUser(id)

	if err != nil {
		response := helpers.NewResponse(data, 400, true, err.Error())
		return response, err
	}

	// if *&data.Model.ID == 0 {
	// 	response := helpers.NewResponse(data, 404, false, "User Not Found")
	// 	return response, nil
	// }

	response := helpers.NewResponse(data, 200, false, "User Found")
	return response, nil
}

func (r *UserService) AddUser(user *models.User) (*helpers.Response, error) {
	hassPssword, err := helpers.HashPasword(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = hassPssword
	data, err := r.repository.AddUser(user)

	if err != nil {
		response := helpers.NewResponse(data, 400, true, err.Error())
		return response, err
	}

	// if *&data.Model.ID == 0 {
	// 	response := helpers.NewResponse(data, 404, false, "User Not Found")
	// 	return response, nil
	// }

	response := helpers.NewResponse(data, 200, false, "Success Add User")
	return response, nil
}

func (r *UserService) UpdateUser(user *models.User) (*helpers.Response, error) {
	data, err := r.repository.UpdateUser(user)

	if err != nil {
		response := helpers.NewResponse(data, 400, true, err.Error())
		return response, err
	}

	// if *&data.Model.ID == 0 {
	// 	response := helpers.NewResponse(data, 404, false, "User Not Found")
	// 	return response, nil
	// }

	response := helpers.NewResponse(data, 200, false, "User Found")
	return response, nil
}

func (r *UserService) DeleteUser(user *models.User) (*helpers.Response, error) {
	err := r.repository.DeleteUser(user)

	if err != nil {
		response := helpers.NewResponse(err, 400, true, err.Error())
		return response, err
	}

	// if *&data.Model.ID == 0 {
	// 	response := helpers.NewResponse(data, 404, false, "User Not Found")
	// 	return response, nil
	// }

	response := helpers.NewResponse(err, 200, false, "Success Delete User")
	return response, nil
}
