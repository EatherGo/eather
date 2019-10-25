package lib

import (
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

var (
	router    *mux.Router
	onceRoute sync.Once
	mr        = mux.NewRouter()
)

type eatherRouter mux.Router

// Routes struct - collection of routes
type Routes struct {
	Collection map[string]func(w http.ResponseWriter, r *http.Request) EatherResponse
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
	http.Handle("/", router)
}
