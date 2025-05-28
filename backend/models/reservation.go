package models

type Reservation struct {
    ID             uint   `json:"id" gorm:"primaryKey"`
    NameOfReserver string `json:"nameofreserver"`
    Date           string `json:"date"` 
    Class          string `json:"class"`
    HotelID        uint   `json:"hotel_id"`
}
