package application

import (
	"API/src/domain/repositories"
    "errors"
)

type UpdateOrder struct {
    db repositories.IntrOrder
}

func NewUpdateOrder(db repositories.IntrOrder) *UpdateOrder {
    return &UpdateOrder{db: db}
}

func (uo *UpdateOrder) Execute(id int32, userID int32, amount float32, status string) error {
    order, err := uo.db.GetByID(id)
    if err != nil {
        return err
    }
    if order == nil {
        return errors.New("order not found")
    }
    order.UpdateOrder(id, userID, float32(amount), status)
    return uo.db.Update(order)
}