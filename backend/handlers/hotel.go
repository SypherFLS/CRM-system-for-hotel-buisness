package handlers

import (
    "net/http"
    // "strconv"
    "github.com/gin-gonic/gin"
    "backend/db"
    "backend/models"
)


func GetHotels(c *gin.Context) {
    var hotels []models.Hotel
    db.DB.Find(&hotels)
    c.JSON(http.StatusOK, hotels)
}


func GetHotel(c *gin.Context) {
    id := c.Param("id")
    var hotel models.Hotel
    if err := db.DB.First(&hotel, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Hotel not found"})
        return
    }
    c.JSON(http.StatusOK, hotel)
}


func CreateHotel(c *gin.Context) {
    var hotel models.Hotel
    if err := c.ShouldBindJSON(&hotel); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if hotel.Picture == "" {
        hotel.Picture = "hotel_futazh.png"
    }
    db.DB.Create(&hotel)
    c.JSON(http.StatusOK, hotel)
}


func UpdateHotel(c *gin.Context) {
    id := c.Param("id")
    var hotel models.Hotel
    if err := db.DB.First(&hotel, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Hotel not found"})
        return
    }
    if err := c.ShouldBindJSON(&hotel); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    db.DB.Save(&hotel)
    c.JSON(http.StatusOK, hotel)
}


func DeleteHotel(c *gin.Context) {
    id := c.Param("id")
    var hotel models.Hotel
    if err := db.DB.First(&hotel, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Hotel not found"})
        return
    }
    db.DB.Delete(&hotel)
    c.JSON(http.StatusOK, gin.H{"message": "Hotel deleted"})
}
