package repository

import (
    "database/sql"
    "project/models"
)

type UserRepo struct {
    DB *sql.DB
}

func (r *UserRepo) GetByLogin(login string) (*models.User, error) {
    user := &models.User{}
    err := r.DB.QueryRow("SELECT id, login, email, password, permission FROM users WHERE login=$1", login).
        Scan(&user.ID, &user.Login, &user.Email, &user.Password, &user.Permission)
    if err != nil {
        return nil, err
    }
    return user, nil
}

// Другие методы: CreateUser, GetByID и т.п.
