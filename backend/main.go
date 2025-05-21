package main

import (
    "log"
    "net/http"
    "project/config"
    "project/handlers"
    "project/middleware"
    "github.com/gorilla/mux"
)

func main() {
    cfg := config.LoadConfig() // Загрузка конфигурации (БД и т.п.)

    r := mux.NewRouter()

    // Роуты авторизации
    r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
    r.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")

    // Роуты для работы со второй таблицей, защищены JWT middleware
    api := r.PathPrefix("/api").Subrouter()
    api.Use(middleware.JWTAuthMiddleware)
    api.HandleFunc("/userdates", handlers.GetUserDates).Methods("GET")
    api.HandleFunc("/userdates", handlers.CreateUserDate).Methods("POST")
    api.HandleFunc("/userdates/{id}", handlers.UpdateUserDate).Methods("PUT")
    api.HandleFunc("/userdates/{id}", handlers.DeleteUserDate).Methods("DELETE")

    log.Println("Server started at :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
