package application

import (	
	"API/src/Users/domain/entities"
	"API/src/Users/domain/repositories"
)

type CreateUser struct {
	db repositories.IntrUser
}

func NewCreateUser(db repositories.IntrUser) *CreateUser {
	return &CreateUser{db: db}
}

func (cu *CreateUser) Execute(name string, email string) error {
	// para crear una nueva instncia
	newUser := entities.NewUser(name, email)
	// se guarda el usuario
	return cu.db.SetUser(newUser)
}