package models

import (
	"eather/lib"
	"eather/src/Modules/AdminUsers/utils"
	"log"

	"github.com/jinzhu/gorm"
)

// AdminUser main struct
type AdminUser struct {
	gorm.Model
	Username string
	Password string
	Email    string
	Role     AdminRole
	RoleID   uint `gorm:"foreignkey:RoleRefer"`
}

// CreateAdmin will create admin user
func CreateAdmin(password string, email string, role string) map[string]bool {
	hash, err := utils.GenerateFromPassword(password)
	if err != nil {
		log.Println(err)
		return map[string]bool{"created": false}
	}

	lib.GetDb().Create(&AdminUser{Username: email, Password: hash, Email: email, RoleID: 1})

	return map[string]bool{"created": true}
}
