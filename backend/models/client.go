package models

import (
    "gorm.io/gorm"
    "time"
)

type Client struct {
    gorm.Model
    FullName       string    `json:"full_name"`
    PassportNumber string    `json:"passport_number"`
    CityFrom       string    `json:"city_from"`
    RoomID         uint      `json:"room_id"`
    PlaceNumber    int       `json:"place_number"`
    ArrivalDate    time.Time `json:"arrival_date"`
    PaymentDays    int       `json:"payment_days"`
    Room           Room      `gorm:"foreignKey:RoomID" json:"-"`
}

func (c *Client) TotalPayment(roomPrice float64) float64 {
    pricePerPlace := roomPrice / float64(c.PlaceNumber)
    return pricePerPlace * float64(c.PaymentDays)
}

func (Client) TableName() string {
    return "client"
}