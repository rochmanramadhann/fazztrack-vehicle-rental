package interfaces

import (
	"github.com/rochmanramadhann/fazztrack-vehicle/src/database/gorm/models"
	"github.com/rochmanramadhann/fazztrack-vehicle/src/helpers"
)

type AuthService interface {
	Login(user *models.User) (*helpers.Response, error)
}
