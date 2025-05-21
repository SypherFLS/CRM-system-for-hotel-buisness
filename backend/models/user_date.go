package models

import "time"

type UserDate struct {
    ID       int64     `db:"id"`
    UserID   int64     `db:"user_id"`
    Date     time.Time `db:"date"`
    Number   int       `db:"number"` // Число, которое нужно хранить и изменять
}
