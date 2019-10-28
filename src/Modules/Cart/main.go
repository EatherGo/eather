package main

import (
	"eather/lib"
	"eather/lib/types"
	"eather/src/Modules/Cart/controller"
	"eather/src/Modules/Cart/models"
)

type module struct{}

var db = lib.GetDb()

func main() {
	// Blank
}

func (m module) MapRoutes() {
	router := lib.GetRouter()

	router.HandleFunc("/add-to-cart", controller.AddToCart).Methods("POST")
}

func (m module) Install() {
	db.AutoMigrate(&models.Cart{}, &models.CartItem{})
	db.Model(&models.CartItem{}).AddForeignKey("cart_id", "carts(id)", "RESTRICT", "RESTRICT").AddForeignKey("product_id", "products(id)", "RESTRICT", "RESTRICT")
}

func (m module) Upgrade(version string) {

}

func (m module) GetEventFuncs() map[string]types.EventFunc {
	return make(map[string]types.EventFunc)
}

// Cart to export
func Cart() (f types.Module, err error) {
	f = module{}
	return
}
