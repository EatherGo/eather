package controller

import (
	"encoding/json"
	"net/http"
	"project/lib"
	"project/src/Modules/Cart/models"
	pmodels "project/src/Modules/Product/models"
)

var db = lib.GetDb()

// AddToCart route
func AddToCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	product := pmodels.Product{}

	productCode := r.FormValue("product_code")
	if len(productCode) == 0 {
		// return lib.NewResponse("error")
		json.NewEncoder(w).Encode(map[string]string{"error": "No product code"})
		return
	}

	db.Select("price, code, id").Where("code = ?", productCode).Find(&product)

	cartID := r.FormValue("cart")
	cart := &models.Cart{}

	if len(cartID) == 0 {
		cart = &models.Cart{User: "Jakub", Price: 0}
		db.Create(&cart)
	} else {
		db.Select("id").Where("id = ?", cartID).First(&cart)
	}

	cartItem := models.CartItem{Qty: 1, Price: product.Price, CartID: cart.ID, ProductID: product.ID}
	db.Create(&cartItem)

	json.NewEncoder(w).Encode(cart)
}
