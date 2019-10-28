package models

import "github.com/jinzhu/gorm"

// AdminRole main struct
type AdminRole struct {
	gorm.Model
	Name string
}
