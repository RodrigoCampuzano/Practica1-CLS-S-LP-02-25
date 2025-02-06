package application

import "API/src/Users/domain/repositories"

type DeleteAllUser struct {
	 db repositories.IntrUser
}

func NewDeleteAllUser( db repositories.IntrUser) *DeleteAllUser {
	return &DeleteAllUser{db: db}
}

func (du *DeleteAllUser) Execute() error{
	return du.db.DeleteAll()
}