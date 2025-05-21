package models

type User struct {
    ID       int64  `db:"id"`
    Login    string `db:"login"`
    Email    string `db:"email"`
    Password string `db:"password"` // Хеш пароля
    Permission string `db:"permission"`
}
