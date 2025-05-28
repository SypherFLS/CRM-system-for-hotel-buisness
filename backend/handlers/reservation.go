package handlers

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
    "backend/db"
    "backend/models"
)


func GetReservations(c *gin.Context) {
    var reservations []models.Reservation
    db.DB.Find(&reservations)
    c.JSON(http.StatusOK, reservations)
}


func GetReservation(c *gin.Context) {
    id := c.Param("id")
    var reservation models.Reservation
    if err := db.DB.First(&reservation, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Reservation not found"})
        return
    }
    c.JSON(http.StatusOK, reservation)
}


func isOverlap(date1, date2 string) bool {
    var start1, end1, start2, end2 int
    fmt.Sscanf(date1, "%d:%d", &start1, &end1)
    fmt.Sscanf(date2, "%d:%d", &start2, &end2)
    return start1 <= end2 && start2 <= end1
}

func CreateReservation(c *gin.Context) {
    var reservation models.Reservation
    if err := c.ShouldBindJSON(&reservation); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    var reservations []models.Reservation
    db.DB.Where("hotel_id = ?", reservation.HotelID).Find(&reservations)
    for _, r := range reservations {
        if isOverlap(r.Date, reservation.Date) {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Выбранный период уже занят!"})
            return
        }
    }
    db.DB.Create(&reservation)
    c.JSON(http.StatusOK, reservation)
}

func UpdateReservation(c *gin.Context) {
    id := c.Param("id")
    var reservation models.Reservation
    if err := db.DB.First(&reservation, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Reservation not found"})
        return
    }
    if err := c.ShouldBindJSON(&reservation); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    db.DB.Save(&reservation)
    c.JSON(http.StatusOK, reservation)
}

// Удалить бронь
func DeleteReservation(c *gin.Context) {
    id := c.Param("id")
    var reservation models.Reservation
    if err := db.DB.First(&reservation, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Reservation not found"})
        return
    }
    db.DB.Delete(&reservation)
    c.JSON(http.StatusOK, gin.H{"message": "Reservation deleted"})
}
