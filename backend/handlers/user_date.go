package handlers

import (
    "encoding/json"
    "net/http"
    "project/services"
    "github.com/gorilla/mux"
    "strconv"
)

var userDateService = &services.UserDateService{ /* инициализация UserDateRepo */ }

func GetUserDates(w http.ResponseWriter, r *http.Request) {
    userID := utils.GetUserIDFromContext(r.Context()) // Извлечь из JWT

    dates, err := userDateService.GetUserDates(userID)
    if err != nil {
        http.Error(w, "Error fetching data", http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(dates)
}

func CreateUserDate(w http.ResponseWriter, r *http.Request) {
    userID := utils.GetUserIDFromContext(r.Context())

    var ud struct {
        Date   string `json:"date"`
        Number int    `json:"number"`
    }
    if err := json.NewDecoder(r.Body).Decode(&ud); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    // Парсинг даты и создание модели
    dateParsed, err := time.Parse("2006-01-02", ud.Date)
    if err != nil {
        http.Error(w, "Invalid date format", http.StatusBadRequest)
        return
    }

    userDate := models.UserDate{
        UserID: userID,
        Date:   dateParsed,
        Number: ud.Number,
    }
    err = userDateService.CreateUserDate(&userDate)
    if err != nil {
        http.Error(w, "Error saving data", http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
}

func UpdateUserDate(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.ParseInt(vars["id"], 10, 64)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    var ud struct {
        Date   string `json:"date"`
        Number int    `json:"number"`
    }
    if err := json.NewDecoder(r.Body).Decode(&ud); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    dateParsed, err := time.Parse("2006-01-02", ud.Date)
    if err != nil {
        http.Error(w, "Invalid date format", http.StatusBadRequest)
        return
    }

    userDate := models.UserDate{
        ID:     id,
        Date:   dateParsed,
        Number: ud.Number,
    }
    // Здесь нужно добавить метод Update в сервис и репозиторий
    err = userDateService.UpdateUserDate(&userDate)
    if err != nil {
        http.Error(w, "Error updating data", http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}

func DeleteUserDate(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.ParseInt(vars["id"], 10, 64)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }
    err = userDateService.DeleteUserDate(id)
    if err != nil {
        http.Error(w, "Error deleting data", http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}
