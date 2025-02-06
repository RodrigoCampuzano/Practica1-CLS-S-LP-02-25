package repositories

import "API/src/Orders/domain/entities"

type IntrOrder interface {
    SetOrder(order *entities.Order) error
    GetAll() ([]*entities.Order, error)
    GetByID(id int32) (*entities.Order, error)
    Update(order *entities.Order) error
    DeleteAll() error
    DeleteByID(id int32) error
}