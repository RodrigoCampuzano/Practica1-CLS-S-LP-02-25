package repositories

import "API/src/Users/domain/entities"

type IntrUser interface {
    SetUser(user *entities.User) error
    GetAll() ([]*entities.User, error)
    GetByID(id int32) (*entities.User, error)
    Update(user *entities.User) error
    DeleteAll() error
    DeleteByID(id int32) error
}