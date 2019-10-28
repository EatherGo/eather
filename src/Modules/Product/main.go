package main

import (
	"eather/lib"
	"eather/lib/types"

	"eather/src/Modules/Product/controller"
	"eather/src/Modules/Product/models"
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
	router.HandleFunc("/delete/{product_code}", controller.Delete).Methods("POST")
}

func (m module) Install() {
	db.AutoMigrate(&models.Product{})
}

func (m module) Upgrade(version string) {

}

func (m module) GetEventFuncs() map[string]types.EventFunc {
	return eventFuncs
}

// Product to export
func Product() (f types.Module, err error) {
	f = module{}
	return
}
