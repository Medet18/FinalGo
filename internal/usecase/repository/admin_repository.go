package repository

import (
	// "finals/internal/domain/model"
	"net/http"
)

type AdminRepository interface {
	Signin(w http.ResponseWriter, r *http.Request)
	SiteAuto(w http.ResponseWriter, r *http.Request)
}
