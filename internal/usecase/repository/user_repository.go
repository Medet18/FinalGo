package repository

import "finals/domain/model"

type UserRepository interface {
	DeleteUsers(u *model.Users) (uuid string)
	CreateNewUsers(u *model.Users) (uuid, name, email, phone, cv_path string)
}
