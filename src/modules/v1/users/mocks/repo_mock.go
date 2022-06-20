package mocks

import (
	"github.com/rochmanramadhann/fazztrack-vehicle/src/database/gorm/models"
	"github.com/stretchr/testify/mock"
)

type RepoMock struct {
	Mock mock.Mock
}

func (pr *RepoMock) SearchUsers(user *models.User) (*models.Users, error) {
	args := pr.Mock.Called(user)
	return args.Get(0).(*models.Users), args.Error(1)
}

func (pr *RepoMock) SortUsers(sort string) (*models.Users, error) {
	args := pr.Mock.Called(sort)
	return args.Get(0).(*models.Users), args.Error(1)
}

func (pr *RepoMock) GetUsers() (*models.Users, error) {
	args := pr.Mock.Called()
	return args.Get(0).(*models.Users), args.Error(1)
}

func (pr *RepoMock) GetUserByUsername(username string) (*models.User, error) {
	args := pr.Mock.Called(username)
	return args.Get(0).(*models.User), args.Error(1)
}

func (pr *RepoMock) GetUser(id int) (*models.User, error) {
	args := pr.Mock.Called(id)
	return args.Get(0).(*models.User), args.Error(1)
}

func (pr *RepoMock) AddUser(user *models.User) (*models.User, error) {
	args := pr.Mock.Called(user)
	return args.Get(0).(*models.User), args.Error(1)
}

func (pr *RepoMock) UpdateUser(user *models.User) (*models.User, error) {
	args := pr.Mock.Called(user)
	return args.Get(0).(*models.User), args.Error(1)
}

func (pr *RepoMock) DeleteUser(user *models.User) error {
	args := pr.Mock.Called(user)
	return args.Error(1)
}
