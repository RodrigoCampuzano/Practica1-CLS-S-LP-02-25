package application

import "API/src/Users/domain/repositories"

type UpdateUser struct {
	db repositories.IntrUser
}

func NewUpdateUser(db repositories.IntrUser) *UpdateUser{
	return &UpdateUser{db: db}
}

func (uu *UpdateUser) Execute(id int32, name string, email string) error {
	updateuser, err := uu.db.GetByID(id)
	if err != nil{
		return err
	}
	updateuser.Name = name
	updateuser.Email = email

	return uu.db.Update(updateuser)
}

