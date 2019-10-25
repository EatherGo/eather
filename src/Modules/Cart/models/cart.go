package models

import (
	pmodels "eather/src/Modules/Product/models"

	"github.com/jinzhu/gorm"
)

// Cart main struct
type Cart struct {
	gorm.Model
	User  string
	Price float64
}

// CartItem main struct
type CartItem struct {
	gorm.Model
	Qty       int
	Price     float64
	Cart      Cart
	CartID    uint `gorm:"foreignkey:CartRefer"`
	Product   pmodels.Product
	ProductID uint `gorm:"foreignkey:ProductRefer"`
}
