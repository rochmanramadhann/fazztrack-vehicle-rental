package middleware

import (
	"net/http"
	"strings"

	"github.com/rochmanramadhann/fazztrack-vehicle/src/helpers"
)

func CheckAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		headerToken := r.Header.Get("Authorization")

		if !strings.Contains(headerToken, "Bearer") {
			helpers.NewResponse("invalid header type", 401, false, "").Send(w)
			return
		}

		token := strings.Replace(headerToken, "Bearer ", "", -1)
		checkToken, err := helpers.CheckToken(token)
		if err != nil {
			helpers.NewResponse(err.Error(), 201, true, "").Send(w)
			return
		}

		if !checkToken {
			helpers.NewResponse("Silahkan login kembali", 401, false, "").Send(w)
			return
		}

		next.ServeHTTP(w, r)
	}
}
