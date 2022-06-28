package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/newrelic/go-agent/v3/integrations/nrgorilla"
	"github.com/newrelic/go-agent/v3/newrelic"
	database "github.com/rochmanramadhann/fazztrack-vehicle/src/database/gorm"
	"github.com/rochmanramadhann/fazztrack-vehicle/src/modules/v1/auth"
	"github.com/rochmanramadhann/fazztrack-vehicle/src/modules/v1/favorites"
	"github.com/rochmanramadhann/fazztrack-vehicle/src/modules/v1/orders"
	"github.com/rochmanramadhann/fazztrack-vehicle/src/modules/v1/users"
	"github.com/rochmanramadhann/fazztrack-vehicle/src/modules/v1/vehicle_types"
	"github.com/rochmanramadhann/fazztrack-vehicle/src/modules/v1/vehicles"
	"github.com/rs/cors"
)

func New() (http.Handler, error) {
	mainRoute := mux.NewRouter()
	nRelic, err := newrelic.NewApplication(
		newrelic.ConfigAppName("fazztrack-vehicle"),
		newrelic.ConfigLicense("7bad166af83620e844b6418ccff6b402996bNRAL"),
		newrelic.ConfigDistributedTracerEnabled(true),
	)
	if err != nil {
		return nil, err
	}

	mainRoute.Use(nrgorilla.Middleware(nRelic))

	db, err := database.New()
	if err != nil {
		return nil, err
	}

	mainRoute.HandleFunc("/", sampleHandler).Methods("GET")
	mainRoute.HandleFunc(newrelic.WrapHandleFunc(nRelic, "/relic", relicHandler)).Methods("GET")
	users.New(mainRoute, db)
	vehicle_types.New(mainRoute, db)
	vehicles.New(mainRoute, db)
	orders.New(mainRoute, db)
	favorites.New(mainRoute, db)
	auth.New(mainRoute, db)

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"OPTIONS", "GET", "POST", "PUT"},
		AllowedHeaders: []string{"Content-Type", "X-CSRF-Token"},
		Debug:          true,
	}).Handler(mainRoute)

	return corsMiddleware, nil
}

func sampleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte("hello worlds"))
}

func relicHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte("hello newrelic"))
}
