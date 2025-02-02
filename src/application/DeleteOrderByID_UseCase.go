package application

import "API/src/domain/repositories"

type DeleteOrder struct {
    db repositories.IntrOrder
}

func NewDeleteOrderByID(db repositories.IntrOrder) *DeleteOrder {
    return &DeleteOrder{db: db}
}

func (do *DeleteOrder) Execute(id int32) error {
    return do.db.DeleteByID(id)
}