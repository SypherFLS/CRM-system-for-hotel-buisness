package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "strings"
    "backend/models"
    "gorm.io/gorm"
)

type EmployeeHandler struct {
    DB *gorm.DB
}

func (h *EmployeeHandler) GetAllEmployees(w http.ResponseWriter, r *http.Request) {
    var employees []models.Employee
    if err := h.DB.Find(&employees).Error; err != nil {
        http.Error(w, "DB error", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    json.NewEncoder(w).Encode(employees)
}

func (h *EmployeeHandler) AddEmployee(w http.ResponseWriter, r *http.Request) {
    var employee models.Employee
    if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }
    if err := h.DB.Create(&employee).Error; err != nil {
        http.Error(w, "DB error", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    json.NewEncoder(w).Encode(employee)
}

func (h *EmployeeHandler) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
    idStr := strings.TrimPrefix(r.URL.Path, "/api/employees/")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }
    var employee models.Employee
    if err := h.DB.First(&employee, id).Error; err != nil {
        http.Error(w, "Employee not found", http.StatusNotFound)
        return
    }
    var updateData struct {
        FullName string `json:"full_name"`
        Floors   string `json:"floors"`
        WeekDays string `json:"week_days"`
    }
    if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }
    employee.FullName = updateData.FullName
    employee.Floors = updateData.Floors
    employee.WeekDays = updateData.WeekDays
    if err := h.DB.Save(&employee).Error; err != nil {
        http.Error(w, "DB error", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    json.NewEncoder(w).Encode(employee)
}

func (h *EmployeeHandler) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.Atoi(r.URL.Query().Get("id"))
    if err := h.DB.Delete(&models.Employee{}, id).Error; err != nil {
        http.Error(w, "DB error", http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}
