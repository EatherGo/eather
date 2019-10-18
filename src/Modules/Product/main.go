package main

import (
	"project/lib"
	"project/lib/interfaces"

	"project/src/Modules/Product/controller"
	"project/src/Modules/Product/models"
)

type module struct{}

var db = lib.GetDb()

func main() {
	// Blank
}

func (m module) MapRoutes() {
	router := lib.GetRouter()

	router.HandleFunc("/products", controller.Index).Methods("GET")
	// router.HandleFunc("products/", controller.Get).Methods("GET")
	router.HandleFunc("/products/store", controller.Store).Methods("POST")
	router.HandleFunc("/delete/", controller.Delete).Methods("POST")
}

func (m module) Install() {
	db.AutoMigrate(&models.Product{})
}

func (m module) Upgrade(version string) {

}

func (m module) GetEventFuncs() map[string]func() {
	return eventFuncs
}

// Product to export
func Product() (f interfaces.Module, err error) {
	f = module{}
	return
}
