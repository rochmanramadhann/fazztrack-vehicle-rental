package interfaces

import (
	"github.com/rochmanramadhann/fazztrack-vehicle/src/database/gorm/models"
	"github.com/rochmanramadhann/fazztrack-vehicle/src/helpers"
)

type UserRepository interface {
	SearchUsers(user *models.User) (*models.Users, error)
	SortUsers(sort string) (*models.Users, error)
	GetUsers() (*models.Users, error)
	GetUser(id int) (*models.User, error)
	AddUser(user *models.User) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
	DeleteUser(user *models.User) error
	GetUserByUsername(username string) (*models.User, error)
}

type UserService interface {
	SearchUsers(user *models.User) (*helpers.Response, error)
	SortUsers(sort string) (*helpers.Response, error)
	GetUsers() (*helpers.Response, error)
	GetUser(id int) (*helpers.Response, error)
	AddUser(user *models.User) (*helpers.Response, error)
	UpdateUser(user *models.User) (*helpers.Response, error)
	DeleteUser(user *models.User) (*helpers.Response, error)
	GetUserByUsername(username string) (*helpers.Response, error)
}
