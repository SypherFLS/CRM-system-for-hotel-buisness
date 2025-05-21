package repository

import (
    "database/sql"
    "project/models"
    "time"
)

type UserDateRepo struct {
    DB *sql.DB
}

func (r *UserDateRepo) GetByUserID(userID int64) ([]models.UserDate, error) {
    rows, err := r.DB.Query("SELECT id, user_id, date, number FROM user_dates WHERE user_id=$1", userID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var results []models.UserDate
    for rows.Next() {
        var ud models.UserDate
        err := rows.Scan(&ud.ID, &ud.UserID, &ud.Date, &ud.Number)
        if err != nil {
            return nil, err
        }
        results = append(results, ud)
    }
    return results, nil
}

func (r *UserDateRepo) Create(ud *models.UserDate) error {
    _, err := r.DB.Exec("INSERT INTO user_dates (user_id, date, number) VALUES ($1, $2, $3)", ud.UserID, ud.Date, ud.Number)
    return err
}

// Аналогично Update и Delete
