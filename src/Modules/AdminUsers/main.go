package main

import (
	"eather/lib"
	"eather/lib/types"
	"eather/src/Modules/AdminUsers/controller"
	"eather/src/Modules/AdminUsers/models"
)

type module struct{}

var db = lib.GetDb()
var mySigningKey = []byte("secret")

func main() {
	// Blank
}

// var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
// 	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
// 		return mySigningKey, nil
// 	},
// 	SigningMethod: jwt.SigningMethodHS256,
// })

func (m module) MapRoutes() {
	router := lib.GetRouter()

	s := router.PathPrefix("/admin").Subrouter()
	s.Use(JwtVerify)

	router.HandleFunc("/admin/login", controller.Login).Methods("POST")
	s.HandleFunc("/create", controller.Create).Methods("POST")
	s.HandleFunc("/me", controller.Me).Methods("GET")
	// router.HandleFunc("/get-token", controller.GetToken).Methods("GET")
}

func (m module) Install() {
	db.AutoMigrate(&models.AdminUser{}, &models.AdminRole{})
}

func (m module) Upgrade(version string) {

}

func (m module) GetEventFuncs() map[string]types.EventFunc {
	return make(map[string]types.EventFunc)
}

// AdminUsers callable func
func AdminUsers() (f types.Module, err error) {
	f = module{}
	return
}
