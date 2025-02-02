package persistence

import (
    "API/src/domain/entities"
    "database/sql"
)

type OrderRepository struct {
    db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
    return &OrderRepository{db: db}
}

func (repo *OrderRepository) SetOrder(order *entities.Order) error {
    _, err := repo.db.Exec("INSERT INTO orders (user_id, amount, status) VALUES (?, ?, ?)", order.UserID, order.Amount, order.Status)
    return err
}

func (repo *OrderRepository) GetAll() ([]*entities.Order, error) {
    rows, err := repo.db.Query("SELECT id, user_id, amount, status FROM orders")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var orders []*entities.Order
    for rows.Next() {
        var order entities.Order
        if err := rows.Scan(&order.ID, &order.UserID, &order.Amount, &order.Status); err != nil {
            return nil, err
        }
        orders = append(orders, &order)
    }
    return orders, nil
}

func (repo *OrderRepository) GetByID(id int32) (*entities.Order, error) {
    row := repo.db.QueryRow("SELECT id, user_id, amount, status FROM orders WHERE id = ?", id)
    var order entities.Order
    if err := row.Scan(&order.ID, &order.UserID, &order.Amount, &order.Status); err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        return nil, err
    }
    return &order, nil
}

func (repo *OrderRepository) Update(order *entities.Order) error {
    _, err := repo.db.Exec("UPDATE orders SET user_id = ?, amount = ?, status = ? WHERE id = ?", order.UserID, order.Amount, order.Status, order.ID)
    return err
}

func (repo *OrderRepository) DeleteAll() error {
    _, err := repo.db.Exec("DELETE FROM orders")
    return err
}

func (repo *OrderRepository) DeleteByID(id int32) error {
    _, err := repo.db.Exec("DELETE FROM orders WHERE id = ?", id)
    return err
}