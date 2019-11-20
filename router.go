package eather

import (
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var (
	router    *mux.Router
	onceRoute sync.Once
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
func RegisterRoutes(corsOpts *cors.Cors) {
	http.Handle("/", corsOpts.Handler(router))
}
