package model

import (
	"encoding/json"

	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("secret_key")

type Admin struct {
	AUID         string
	Privilege_ID int    `json:"pr"`
	IsActive     bool   `json:"isactive"`
	Password     string `json:"password"`
	Username     string `json:"username"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// func (a *Admin) isChecked(login, password string) bool {
// 	if a.Login == login && a.Password == password {
// 		return true
// 	}
// 	return false
// }
var users = map[string]string{}

func Signin(w http.ResponseWriter, r *http.Request) {

	var admins Admin
	// Получить тело JSON и декодировать в учетные данные
	error := json.NewDecoder(r.Body).Decode(&admins)
	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	expectedPassword, ok := users[admins.Username]
	if !ok || expectedPassword != admins.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	expirationTime := time.Now().Add(30 * time.Minute)
	claims := &Claims{
		Username: admins.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(jwtKey)

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

func SiteAuto(w http.ResponseWriter, r *http.Request) {

	c, err := r.Cookie("token")
	if err == http.ErrNoCookie {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	tknStr := c.Value
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Username)))
}
