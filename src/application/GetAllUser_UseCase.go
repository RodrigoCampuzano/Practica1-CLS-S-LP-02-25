package application

import (
	"API/src/domain/entities"
	"API/src/domain/repositories"
)

type GetUsers struct {
	db repositories.IntrUser
}

func NewGetUser(db repositories.IntrUser) *GetUsers {
	return &GetUsers{db: db}
}

func (vu *GetUsers) Execute() ([]*entities.User, error){
	return vu.db.GetAll()
}
