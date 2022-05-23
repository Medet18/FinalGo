package presenter

import "finals/internal/domain/model"

type adminPresenter struct {
}

type AdminPresenter interface {
	ResponseAdmins(admins []*model.Admin) []*model.Admin
}

func newAdminPresenter() AdminPresenter {
	return &adminPresenter{}
}

func (ap *adminPresenter) ResponseAdmins(admins []*model.Admin) []*model.Admin {
	for _, u := range admins {
		u.Username = "Admins name : " + u.Username
	}
	return admins
}
