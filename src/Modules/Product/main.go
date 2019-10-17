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

	router.AddGet("products", controller.Index)
	router.AddGet("products/", controller.Get)
	router.AddPost("products/store", controller.Store)
	router.AddGet("delete/", controller.Delete)
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
