package models

import "github.com/jinzhu/gorm"

// Product main struct
type Product struct {
	gorm.Model
	Code  string `gorm:"unique;not null"`
	Price float64
}
