package handlers

import (
    "encoding/json"
    "net/http"
    "project/services"
    "project/utils"
)

var authService = &services.AuthService{ /* инициализация UserRepo */ }

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    var creds struct {
        Login    string `json:"login"`
        Password string `json:"password"`
    }
    if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    user, err := authService.Authenticate(creds.Login, creds.Password)
    if err != nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    token, err := utils.GenerateJWT(user.ID, user.Permission)
    if err != nil {
        http.Error(w, "Failed to generate token", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{"token": token})
}
