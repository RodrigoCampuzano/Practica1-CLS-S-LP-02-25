package application

import (
	"API/src/Users/domain/repositories"
	"API/src/Users/domain/entities"
)

type GetUserByID struct {
    db repositories.IntrUser
}

func NewGetUserByID(db repositories.IntrUser) *GetUserByID {
    return &GetUserByID{db: db}
}

func (gub *GetUserByID) Execute(id int32) (*entities.User, error) {
    return gub.db.GetByID(id)
}