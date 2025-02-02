package persistence

import (
    "API/src/domain/entities"
    "database/sql"
)

type UserRepository struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
    return &UserRepository{db: db}
}

func (repo *UserRepository) SetUser(user *entities.User) error {
    _, err := repo.db.Exec("INSERT INTO user (name, email) VALUES (?, ?)", user.Name, user.Email)
    return err
}

func (repo *UserRepository) GetAll() ([]*entities.User, error) {
    rows, err := repo.db.Query("SELECT id, name, email FROM user")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []*entities.User
    for rows.Next() {
        var user entities.User
        if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
            return nil, err
        }
        users = append(users, &user)
    }
    return users, nil
}

func (repo *UserRepository) GetByID(id int32) (*entities.User, error) {
    row := repo.db.QueryRow("SELECT id, name, email FROM user WHERE id = ?", id)
    var user entities.User
    if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        return nil, err
    }
    return &user, nil
}

func (repo *UserRepository) Update(user *entities.User) error {
    _, err := repo.db.Exec("UPDATE user SET name = ?, email = ? WHERE id = ?", user.Name, user.Email, user.ID)
    return err
}

func (repo *UserRepository) DeleteAll() error {
    _, err := repo.db.Exec("DELETE FROM user")
    return err
}

func (repo *UserRepository) DeleteByID(id int32) error {
    _, err := repo.db.Exec("DELETE FROM user WHERE id = ?", id)
    return err
}