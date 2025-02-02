package application

import "API/src/domain/repositories"

type DeleteUser struct {
    db repositories.IntrUser
}

func NewDeleteUserByID(db repositories.IntrUser) *DeleteUser {
    return &DeleteUser{db: db}
}

func (du *DeleteUser) Execute(id int32) error {
    return du.db.DeleteByID(id)
}