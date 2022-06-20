package users

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/rochmanramadhann/fazztrack-vehicle/src/database/gorm/models"
	"github.com/rochmanramadhann/fazztrack-vehicle/src/modules/v1/users/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var searchRequest = models.User{
	Email: "oman@mail.com",
}

const sortRequest = "asc"

var byUsernameRequest = models.User{
	Name: "oman",
}

var singleRequest = models.User{
	Model: gorm.Model{ID: 1},
	Name: "oman",
}

var addRequest = models.User{
	Model: gorm.Model{ID: 2},
}

var updateRequest = models.User{
	Model: gorm.Model{ID: 2},
}

var deleteRequest = models.User{
	Model: gorm.Model{ID: 2},
}

func TestSearchUsers(t *testing.T) {
	var repository = mocks.RepoMock{Mock: mock.Mock{}}
	var service = UserService{&repository}

	repository.Mock.On("SearchUsers", &searchRequest).Return(searchRequest, nil)
	data, err := service.SearchUsers(&searchRequest)
	fmt.Println(searchRequest)
	fmt.Println("hahaasdas",reflect.TypeOf(data).Kind())

	users := data.Data.(*models.Users)
	tempEmail := ""
	for _, v := range *users {
		if v.Email == searchRequest.Email {
			tempEmail = v.Email
		}
	}
	assert.Equal(t, "oman@mail.com", tempEmail, "Salah")
	assert.Nil(t, err)
}

func TestSortUsers(t *testing.T) {
	var repository = mocks.RepoMock{Mock: mock.Mock{}}
	var service = UserService{&repository}

	repository.Mock.On("GetUser", 1).Return(&searchRequest, nil)
	data, err := service.GetUser(1)

	users := data.Data.(*models.User)
	assert.Equal(t, "oman", users.Name, "Salah")
	assert.Nil(t, err)
}

func TestGetUsers(t *testing.T) {
	var repository = mocks.RepoMock{Mock: mock.Mock{}}
	var service = UserService{&repository}

	repository.Mock.On("GetUser", 1).Return(&searchRequest, nil)
	data, err := service.GetUser(1)

	users := data.Data.(*models.User)
	assert.Equal(t, "oman", users.Name, "Salah")
	assert.Nil(t, err)
}

func TestGetUserByUsername(t *testing.T) {
	var repository = mocks.RepoMock{Mock: mock.Mock{}}
	var service = UserService{&repository}

	repository.Mock.On("GetUser", 1).Return(&searchRequest, nil)
	data, err := service.GetUser(1)

	users := data.Data.(*models.User)
	assert.Equal(t, "oman", users.Name, "Salah")
	assert.Nil(t, err)
}

func TestGetUser(t *testing.T) {
	var repository = mocks.RepoMock{Mock: mock.Mock{}}
	var service = UserService{&repository}

	repository.Mock.On("GetUser", 1).Return(&singleRequest, nil)
	data, err := service.GetUser(1)

	users := data.Data.(*models.User)
	assert.Equal(t, "oman", users.Name, err)
	assert.Nil(t, err)
}

func TestAddUser(t *testing.T) {
	var repository = mocks.RepoMock{Mock: mock.Mock{}}
	var service = UserService{&repository}

	repository.Mock.On("GetUser", 1).Return(&searchRequest, nil)
	data, err := service.GetUser(1)

	users := data.Data.(*models.User)
	assert.Equal(t, "oman", users.Name, "Salah")
	assert.Nil(t, err)
}

func TestUpdateUser(t *testing.T) {
	var repository = mocks.RepoMock{Mock: mock.Mock{}}
	var service = UserService{&repository}

	repository.Mock.On("GetUser", 1).Return(&searchRequest, nil)
	data, err := service.GetUser(1)

	users := data.Data.(*models.User)
	assert.Equal(t, "oman", users.Name, "Salah")
	assert.Nil(t, err)
}

func TestDeleteUser(t *testing.T) {
	var repository = mocks.RepoMock{Mock: mock.Mock{}}
	var service = UserService{&repository}

	repository.Mock.On("GetUser", 1).Return(&searchRequest, nil)
	data, err := service.GetUser(1)

	users := data.Data.(*models.User)
	assert.Equal(t, "oman", users.Name, "Salah")
	assert.Nil(t, err)
}
