package application

import "API/src/Orders/domain/repositories"

type DeleteAllOrders struct {
    db repositories.IntrOrder
}

func NewDeleteAllOrders(db repositories.IntrOrder) *DeleteAllOrders {
    return &DeleteAllOrders{db: db}
}

func (dao *DeleteAllOrders) Execute() error {
    return dao.db.DeleteAll()
}