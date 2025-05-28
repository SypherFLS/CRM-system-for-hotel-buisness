package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"backend/models"
	"gorm.io/gorm"
)

type CleaningHandler struct {
	DB *gorm.DB
}

func (h *CleaningHandler) GetCleanings(w http.ResponseWriter, r *http.Request) {
	var cleanings []models.CleaningSchedule
	query := h.DB.Model(&models.CleaningSchedule{}).
		Preload("Room").
		Preload("Employee")

	if date := r.URL.Query().Get("date"); date != "" {
		query = query.Where("cleaning_date::date = ?", date)
	}
	if roomID := r.URL.Query().Get("room_id"); roomID != "" {
		query = query.Where("room_id = ?", roomID)
	}
	if employeeID := r.URL.Query().Get("employee_id"); employeeID != "" {
		query = query.Where("employee_id = ?", employeeID)
	}

	if err := query.Find(&cleanings).Error; err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(cleanings)
}

func (h *CleaningHandler) AddCleaning(w http.ResponseWriter, r *http.Request) {
	var cleaning models.CleaningSchedule
	if err := json.NewDecoder(r.Body).Decode(&cleaning); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	
	var existing models.CleaningSchedule
	if err := h.DB.Where("employee_id = ? AND cleaning_date::date = ?", cleaning.EmployeeID, cleaning.CleaningDate.Format("2006-01-02")).First(&existing).Error; err == nil {
		http.Error(w, "Employee already assigned to cleaning on this date", http.StatusConflict)
		return
	}
	if err := h.DB.Create(&cleaning).Error; err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(cleaning)
}

func (h *CleaningHandler) UpdateCleaning(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var cleaning models.CleaningSchedule
	if err := h.DB.First(&cleaning, id).Error; err != nil {
		http.Error(w, "Cleaning not found", http.StatusNotFound)
		return
	}
	var updateData models.CleaningSchedule
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	
	var existing models.CleaningSchedule
	if err := h.DB.Where("employee_id = ? AND cleaning_date::date = ? AND id != ?", updateData.EmployeeID, updateData.CleaningDate.Format("2006-01-02"), id).First(&existing).Error; err == nil {
		http.Error(w, "Employee already assigned to cleaning on this date", http.StatusConflict)
		return
	}
	cleaning.EmployeeID = updateData.EmployeeID
	cleaning.RoomID = updateData.RoomID
	cleaning.CleaningDate = updateData.CleaningDate
	if err := h.DB.Save(&cleaning).Error; err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(cleaning)
}

func (h *CleaningHandler) DeleteCleaning(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if err := h.DB.Delete(&models.CleaningSchedule{}, id).Error; err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
