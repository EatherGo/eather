package controller

import (
	"eather/lib"
	"eather/src/Modules/AdminUsers/models"
	"eather/src/Modules/AdminUsers/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var db = lib.GetDb()
var mySigningKey = []byte("secret")

// Login route
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user := &models.AdminUser{}
	err := json.NewDecoder(r.Body).Decode(user)

	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Invalid request"}
		json.NewEncoder(w).Encode(resp)
		return
	}

	resp := findAdminUser(user.Email, user.Password)

	json.NewEncoder(w).Encode(resp)

	// json.NewEncoder(w).Encode(map[string]string{"error": "Invalid credentials"})
}

// Create new admin user
func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	email := r.FormValue("email")
	password := r.FormValue("password")

	json.NewEncoder(w).Encode(models.CreateAdmin(password, email, "admin"))
}

func Me(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user")
	json.NewEncoder(w).Encode(user)
}

func findAdminUser(email, password string) map[string]interface{} {
	user := models.AdminUser{}

	if err := lib.GetDb().Where("email = ?", email).First(&user).Error; err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Email address not found"}
		return resp
	}

	expiresAt := time.Now().Add(time.Minute * 100000).Unix()

	match, errf := utils.ComparePasswordAndHash(password, user.Password)
	if (errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword) || match == false {
		var resp = map[string]interface{}{"status": false, "message": "Invalid login credentials. Please try again"}
		return resp
	}

	tk := &models.Token{
		UserID:   user.ID,
		Username: user.Username,
		Email:    user.Email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)

	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		fmt.Println(error)
	}

	var resp = map[string]interface{}{"status": false, "message": "logged in"}
	user.Password = ""
	resp["token"] = tokenString //Store the token in the response
	resp["user"] = user
	return resp

}
