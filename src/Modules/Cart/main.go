package main

import (
	"project/lib"
	"project/lib/interfaces"
	"project/src/Modules/Cart/controller"
	"project/src/Modules/Cart/models"
)

type module struct{}

var db = lib.GetDb()

func main() {
	// Blank
}

func (m module) MapRoutes() {
	router := lib.GetRouter()

	router.AddPost("add-to-cart", controller.AddToCart)
}

func (m module) Install() {
	db.AutoMigrate(&models.Cart{}, &models.CartItem{})
	db.Model(&models.CartItem{}).AddForeignKey("cart_id", "carts(id)", "RESTRICT", "RESTRICT").AddForeignKey("product_id", "products(id)", "RESTRICT", "RESTRICT")
}

func (m module) Upgrade(version string) {

}

func (m module) GetEventFuncs() map[string]func() {
	return make(map[string]func())
}

// Cart to export
func Cart() (f interfaces.Module, err error) {
	f = module{}
	return
}
