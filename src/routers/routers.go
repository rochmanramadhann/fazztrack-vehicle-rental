package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	database "github.com/rochmanramadhann/fazztrack-vehicle/src/database/gorm"
	"github.com/rochmanramadhann/fazztrack-vehicle/src/modules/v1/auth"
	"github.com/rochmanramadhann/fazztrack-vehicle/src/modules/v1/favorites"
	"github.com/rochmanramadhann/fazztrack-vehicle/src/modules/v1/orders"
	"github.com/rochmanramadhann/fazztrack-vehicle/src/modules/v1/users"
	"github.com/rochmanramadhann/fazztrack-vehicle/src/modules/v1/vehicle_types"
	"github.com/rochmanramadhann/fazztrack-vehicle/src/modules/v1/vehicles"
)

func New() (*mux.Router, error) {
	mainRoute := mux.NewRouter()

	db, err := database.New()
	if err != nil {
		return nil, err
	}

	mainRoute.HandleFunc("/", sampleHandler).Methods("GET")
	users.New(mainRoute, db)
	vehicle_types.New(mainRoute, db)
	vehicles.New(mainRoute, db)
	orders.New(mainRoute, db)
	favorites.New(mainRoute, db)
	auth.New(mainRoute, db)

	return mainRoute, nil
}

func sampleHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello worlds"))
}
