package interactor

import (
	"encoding/json"
	"finals/internal/domain/model"
	"finals/internal/usecase/presenter"
	"finals/internal/usecase/repository"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

type adminInteractor struct {
	AdminRepository repository.AdminRepository
	AdminPresenter  presenter.AdminPresenter
}

var jwtKey = []byte("secret_key")

var users = map[string]string{}

// Signin implements AdminInteractor
func (us *adminInteractor) Signin(w http.ResponseWriter, r *http.Request) {
	var admins model.Admin
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
	claims := &model.Claims{
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

// SiteAuto implements AdminInteractor
func (*adminInteractor) SiteAuto(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token")
	if err == http.ErrNoCookie {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	tknStr := c.Value
	claims := &model.Claims{}

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

type AdminInteractor interface {
	Signin(w http.ResponseWriter, r *http.Request)
	SiteAuto(w http.ResponseWriter, r *http.Request)
}

func newAdminInteractor(repo repository.AdminRepository, pres presenter.AdminPresenter) AdminInteractor {
	return &adminInteractor{repo, pres}
}
