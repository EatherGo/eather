package models

import (
	"github.com/jinzhu/gorm"
)

// AdminUser main struct
type AdminUser struct {
	gorm.Model
	Username string
	Password string
	Email    string
	Role     string // TODO create roles
}

// CartItem main struct
// type CartItem struct {
// 	gorm.Model
// 	Qty       int
// 	Price     float64
// 	Cart      Cart
// 	CartID    uint `gorm:"foreignkey:CartRefer"`
// 	Product   pmodels.Product
// 	ProductID uint `gorm:"foreignkey:ProductRefer"`
// }
