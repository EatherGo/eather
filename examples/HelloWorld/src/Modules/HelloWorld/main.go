package main

import (
	"net/http"

	"github.com/EatherGo/eather"
)

type module struct{}

func (m module) MapRoutes() {
	router := eather.GetRouter()

	router.HandleFunc("/", controllerHelloWorld).Methods("GET")
}

// HelloWorld to export in plugin
func HelloWorld() (f eather.Module, err error) {
	f = module{}
	return
}

func controllerHelloWorld(w http.ResponseWriter, r *http.Request) {
	eather.SendJSONResponse(w, eather.Response{Message: "Hello world."})
}
