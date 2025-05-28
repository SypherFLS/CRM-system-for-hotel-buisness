package main

import (
    "github.com/gin-gonic/gin"
    "backend/db"
    "backend/handlers"
)


func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(200)
            return
        }

        c.Next()
    }
}

func main() {
    db.Init()
    r := gin.Default()

    r.Use(CORSMiddleware())

    r.GET("/hotels", handlers.GetHotels)
    r.GET("/hotels/:id", handlers.GetHotel)
    r.POST("/hotels", handlers.CreateHotel)
    r.PUT("/hotels/:id", handlers.UpdateHotel)
    r.DELETE("/hotels/:id", handlers.DeleteHotel)

    r.GET("/reservations", handlers.GetReservations)
    r.GET("/reservations/:id", handlers.GetReservation)
    r.POST("/reservations", handlers.CreateReservation)
    r.PUT("/reservations/:id", handlers.UpdateReservation)
    r.DELETE("/reservations/:id", handlers.DeleteReservation)

    r.Run(":8080")
}
