package main

import (
	"eather/lib"
	"eather/lib/interfaces"
	"eather/src/Modules/AdminUsers/models"
)

type module struct{}

var db = lib.GetDb()

func main() {
	// Blank
}

func (m module) MapRoutes() {
	// router := lib.GetRouter()

	// router.AddPost("add-to-cart", controller.AddToCart)
}

func (m module) Install() {
	db.AutoMigrate(&models.AdminUser{})
}

func (m module) Upgrade(version string) {

}

func (m module) GetEventFuncs() map[string]interfaces.EventFunc {
	return make(map[string]interfaces.EventFunc)
}

// AdminUsers callable func
func AdminUsers() (f interfaces.Module, err error) {
	f = module{}
	return
}
