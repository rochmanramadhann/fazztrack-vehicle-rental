package auth

import (
	"github.com/rochmanramadhann/fazztrack-vehicle/src/database/gorm/models"
	"github.com/rochmanramadhann/fazztrack-vehicle/src/helpers"
	"github.com/rochmanramadhann/fazztrack-vehicle/src/interfaces"
)

type TokenResponse struct {
	Tokens string `json:"token"`
}

type AuthService struct {
	repository interfaces.UserRepository
}

func NewAuthService(repository interfaces.UserRepository) *AuthService {
	return &AuthService{repository}
}

func (r *AuthService) Login(body *models.User) (*helpers.Response, error) {
	user, err := r.repository.GetUserByUsername(body.Name)
	if err != nil {
		return nil, err
	}

	if !helpers.CheckPassword(user.Password, body.Password) {
		return helpers.NewResponse(user, 401, false, "Password Salah"), nil
	}

	token := helpers.NewToken(body.Name)
	theToken, err := token.Create()
	if err != nil {
		return nil, err
	}

	response := helpers.NewResponse(TokenResponse{Tokens: theToken}, 200, false, "Login Success")
	return response, nil

}
