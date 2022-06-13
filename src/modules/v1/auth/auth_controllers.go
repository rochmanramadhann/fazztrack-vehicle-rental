package auth

import (
	"encoding/json"
	"net/http"

	"github.com/rochmanramadhann/fazztrack-vehicle/src/database/gorm/models"
	"github.com/rochmanramadhann/fazztrack-vehicle/src/interfaces"
)

type AuthController struct {
	service interfaces.AuthService
}

func NewCtrl(service interfaces.AuthService) *AuthController {
	return &AuthController{service}
}

func (ctrl *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User

	json.NewDecoder(r.Body).Decode(&user)
	result, err := ctrl.service.Login(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	result.Send(w)
}
