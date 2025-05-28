package handlers

import (
    "encoding/json"
    "log"
    "net/http"
    "strconv"
    "strings"
    "backend/models"
    "gorm.io/gorm"
)

type ClientHandler struct {
    DB *gorm.DB
}

func (h *ClientHandler) GetAllClients(w http.ResponseWriter, r *http.Request) {
    var clients []models.Client
    if err := h.DB.Find(&clients).Error; err != nil {
        http.Error(w, "DB error", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    json.NewEncoder(w).Encode(clients)
}

func (h *ClientHandler) GetClientsByCity(w http.ResponseWriter, r *http.Request) {
    city := r.URL.Query().Get("city")
    var clients []models.Client
    if err := h.DB.Where("city_from = ?", city).Find(&clients).Error; err != nil {
        http.Error(w, "DB error", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    json.NewEncoder(w).Encode(clients)
}

func (h *ClientHandler) GetClientsInSingleRooms(w http.ResponseWriter, r *http.Request) {
    var clients []models.Client
    if err := h.DB.Joins("JOIN rooms ON rooms.id = clients.room_id").
        Where("rooms.room_type = ?", "single").
        Find(&clients).Error; err != nil {
        http.Error(w, "DB error", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    json.NewEncoder(w).Encode(clients)
}

func (h *ClientHandler) GetTotalPayment(w http.ResponseWriter, r *http.Request) {
    type Result struct{ Total float64 }
    var result Result
    if err := h.DB.Raw(`
        SELECT SUM(rooms.price_per_day / rooms.capacity * clients.payment_days) as total
        FROM clients
        JOIN rooms ON clients.room_id = rooms.id
    `).Scan(&result).Error; err != nil {
        http.Error(w, "DB error", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    json.NewEncoder(w).Encode(map[string]float64{"total_payment": result.Total})
}

func (h *ClientHandler) DeleteClient(w http.ResponseWriter, r *http.Request) {
    // Извлекаем ID из URL
    idStr := strings.TrimPrefix(r.URL.Path, "/api/clients/")
    idStr = strings.TrimSuffix(idStr, "/")
    log.Println("DeleteClient: idStr =", idStr)
    id, err := strconv.Atoi(idStr)
    if err != nil {
        log.Println("DeleteClient: invalid ID =", idStr)
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }
    var client models.Client
    if err := h.DB.First(&client, id).Error; err != nil {
        log.Println("DeleteClient: client not found, id =", id)
        http.Error(w, "Client not found", http.StatusNotFound)
        return
    }
    if err := h.DB.Delete(&client).Error; err != nil {
        log.Println("DeleteClient: DB error, id =", id)
        http.Error(w, "DB error", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.WriteHeader(http.StatusNoContent)
}
