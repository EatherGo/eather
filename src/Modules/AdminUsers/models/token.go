package models

import "github.com/dgrijalva/jwt-go"

//Token struct declaration
type Token struct {
	UserID   uint
	Username string
	Email    string
	Role     string
	*jwt.StandardClaims
}
