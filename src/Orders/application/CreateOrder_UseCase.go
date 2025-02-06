package application

import (
	"API/src/Orders/domain/entities"
	"API/src/Orders/domain/repositories"
)

type CreateOrder struct {
    db repositories.IntrOrder
}

func NewCreateOrder(db repositories.IntrOrder) *CreateOrder {
    return &CreateOrder{db: db}
}

func (co *CreateOrder) Execute(userID int32, amount float32, status string) error {
    newOrder := entities.NewOrder(userID, amount, status)
    return co.db.SetOrder(newOrder)
}