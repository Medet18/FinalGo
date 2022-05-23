package controller

import (
	"finals/internal/usecase/interactor"
	"net/http"
)

type adminController struct {
	adminInteractor interactor.AdminInteractor
}

type AdminController interface {
	Signin(w http.ResponseWriter, r *http.Request)
	SiteAuto(w http.ResponseWriter, r *http.Request)
}

func newAdminController(adminInteractor interactor.AdminInteractor) AdminController {
	return &adminController{adminInteractor}
}

func (adminController *adminController) Signin(w http.ResponseWriter, r *http.Request) {
	adminController.adminInteractor.Signin(w, r)
}

// SiteAuto implements AdminController
func (adminController *adminController) SiteAuto(w http.ResponseWriter, r *http.Request) {
	adminController.adminInteractor.SiteAuto(w, r)
}
