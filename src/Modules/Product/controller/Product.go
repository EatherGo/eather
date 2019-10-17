package controller

import (
	"fmt"
	"net/http"
	"project/lib"
	"project/src/Modules/Product/models"
	"strconv"
)

var db = lib.GetDb()

// Index route
func Index(w http.ResponseWriter, r *http.Request) lib.EatherResponse {
	products := []models.Product{}

	db.Select("price, code").Find(&products)

	return lib.NewResponse(products)
}

// Get one product
func Get(w http.ResponseWriter, r *http.Request) lib.EatherResponse {
	product := models.Product{}

	id, err := strconv.Atoi(r.URL.Path[len("/products/"):])
	if err != nil {
		fmt.Println(err)
	}

	db.Select("price, code").Where("code = ?", id).First(&product)

	return lib.NewResponse(product)
}

// Store route
func Store(w http.ResponseWriter, r *http.Request) lib.EatherResponse {
	code := r.FormValue("code")
	price := r.FormValue("price")
	priceInt, _ := strconv.ParseFloat(price, 10)

	fmt.Println(code)
	fmt.Println(priceInt)
	product := models.Product{Code: code, Price: priceInt}
	db.Create(&product)

	eventer := lib.GetEvents()

	eventer.Emmit("product_added")
	eventer.Emmit("product_removed")

	return lib.NewResponse(product)
}

// Delete route
func Delete(w http.ResponseWriter, r *http.Request) lib.EatherResponse {
	products := []models.Product{}

	id := r.URL.Path[len("/delete/"):]
	num, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(num)
	}

	db.Select("price, code").Find(&products)

	return lib.NewResponse(products)
}
