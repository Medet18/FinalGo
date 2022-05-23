package presenter

import "finals/internal/domain/model"

type AdminPresenter interface {
	ResponseAdmins(admin *model.Admin) (*model.Admin, error)
}
