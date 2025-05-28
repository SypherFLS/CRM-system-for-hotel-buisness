package models

type Hotel struct {
    ID               uint    `json:"id" gorm:"primaryKey"`
    Name             string  `json:"name" gorm:"column:name"`
    QualityOfNumbers int     `json:"qualityofnumbers" gorm:"column:quality_of_numbers"`
    Location         string  `json:"location" gorm:"column:location"`
    Characteristics  string  `json:"characteristics" gorm:"column:characteristics"`
    PricePerNight    float64 `json:"pricepernight" gorm:"column:price_per_night"`
    Picture          string  `json:"picture" gorm:"column:picture"`
    Reservations     []Reservation `json:"reservations,omitempty"`
}
