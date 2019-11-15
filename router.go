package eather

import (
	"net/http"
	"os"
	"sync"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var (
	router    *mux.Router
	onceRoute sync.Once
	mr        = mux.NewRouter()
)

type eatherRouter mux.Router

// Routes struct - collection of routes
type Routes struct {
	Collection map[string]func(w http.ResponseWriter, r *http.Request) Response
}

// GetRouter - return route collection
func GetRouter() *mux.Router {
	onceRoute.Do(func() {
		router = mux.NewRouter()
	})

	return router
}

// RegisterRoutes - listen for routes
func RegisterRoutes() {
	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{os.Getenv("FRONTEND_URL")}, //you service is available and allowed for this base url
		AllowedMethods: []string{
			http.MethodGet, //http methods for your app
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		},

		AllowedHeaders: []string{
			"*", //or you can your header key values which you are using in your application

		},
	})

	http.Handle("/", corsOpts.Handler(router))
}
