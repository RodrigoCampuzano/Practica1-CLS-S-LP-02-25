package application

import (
	"API/src/Orders/domain/repositories"
	"API/src/Orders/domain/entities"
)

type GetOrderByID struct {
    db repositories.IntrOrder
}

func NewGetOrderByID(db repositories.IntrOrder) *GetOrderByID {
    return &GetOrderByID{db: db}
}

func (gob *GetOrderByID) Execute(id int32) (*entities.Order, error) {
    return gob.db.GetByID(id)
}