package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"project/lib"
	"project/src/Modules/Product/models"
	"strconv"

	"github.com/gorilla/mux"
)

var db = lib.GetDb()

// Index route
func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	products := []models.Product{}

	db.Select("price, code").Find(&products)

	json.NewEncoder(w).Encode(products)
}

// Get one product
func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	product := models.Product{}

	id, err := strconv.Atoi(r.URL.Path[len("/products/"):])
	if err != nil {
		fmt.Println(err)
	}

	db.Select("price, code").Where("code = ?", id).First(&product)

	json.NewEncoder(w).Encode(product)
}

// Store route
func Store(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	fmt.Println(vars)

	code := r.FormValue("code")
	price := r.FormValue("price")
	priceInt, _ := strconv.ParseFloat(price, 10)

	fmt.Println(code)
	fmt.Println(priceInt)
	product := models.Product{Code: code, Price: priceInt}
	if dbr := db.Create(&product); dbr.Error != nil {
		json.NewEncoder(w).Encode(map[string]string{"Error": "Product already exists"})
		return
	}

	lib.GetEvents().Emmit("product_added", map[string]int{"product_code": 50})

	json.NewEncoder(w).Encode(product)
}

// Delete route
func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	db.Select("code").Where("code = ?", vars["product_code"]).Delete(models.Product{})

	lib.GetEvents().Emmit("product_removed", vars["product_code"])

	json.NewEncoder(w).Encode("Product " + vars["product_code"] + " was removed")
}
