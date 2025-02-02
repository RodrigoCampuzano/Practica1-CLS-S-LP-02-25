package application

import (
	"API/src/domain/entities"
	"API/src/domain/repositories"
)

type GetAllOrders struct {
    db repositories.IntrOrder
}

func NewGetAllOrders(db repositories.IntrOrder) *GetAllOrders {
    return &GetAllOrders{db: db}
}

func (gao *GetAllOrders) Execute() ([]*entities.Order, error) {
    return gao.db.GetAll()
}