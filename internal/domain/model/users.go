package model
import 	"fmt"


type Users struct {
	UUID  string 	`gorm:"primary_key" json:"uuid"`
	Name  string	`json:"name"`
	Email string	`json:"email"`
	Phone string	`json:"phone"`
	CV    string 	`json:"cv"`
}

func (u *Users) CreateNewUsers(uuid, name, email, phone, cv_path string) error {
	if uuid == u.UUID || email == u.Email {
		fmt.Errorf("Such a user already exists")
	}
	u.UUID = uuid
	u.Name = name
	u.Email = email
	u.Phone = phone
	u.CV = cv_path
	return nil
}
func(u *Users) DeleteUsers(uuid string) error { 
	if u.UUID == uuid {
		u.Name = ""
		u.Email = ""
		u.Phone = ""
		u.CV = ""
	}
	return nil 
}
 