package models

import (
	pmodels "eather/src/Modules/Product/models"

	"github.com/jinzhu/gorm"
)

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
