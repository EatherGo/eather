package lib

import (
	"encoding/json"
	"net/http"
	"sync"
)

var (
	routes    *Routes
	onceRoute sync.Once
)

// Routes struct - collection of routes
type Routes struct {
	Collection map[string]func(w http.ResponseWriter, r *http.Request) EatherResponse
}

// GetRouter - return route collection
func GetRouter() *Routes {
	onceRoute.Do(func() {
		routes = &Routes{make(map[string]func(w http.ResponseWriter, r *http.Request) EatherResponse)}
	})

	return routes
}

// AddGet - add GET type of route
func (r *Routes) AddGet(name string, f func(w http.ResponseWriter, r *http.Request) EatherResponse) {
	routes.Collection[name] = f
	http.Handle("/"+name, JSONResponse(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}), f, http.MethodGet))
}

// AddPost - add POST type of route
func (r *Routes) AddPost(name string, f func(w http.ResponseWriter, r *http.Request) EatherResponse) {
	routes.Collection[name] = f
	http.Handle("/"+name, JSONResponse(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}), f, http.MethodPost))
}

// JSONResponse to create json response
func JSONResponse(h http.Handler, f func(w http.ResponseWriter, r *http.Request) EatherResponse, method string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if method == r.Method {

			h.ServeHTTP(w, r)

			response := f(w, r)

			js, err := json.Marshal(response.Data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(js)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}

	})
}
