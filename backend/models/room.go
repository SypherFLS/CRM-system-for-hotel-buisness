package models



type Room struct {
    ID          uint    `gorm:"primaryKey" json:"id"`
    Floor       int     `json:"floor"`
    RoomType    string  `json:"room_type"`
    PricePerDay float64 `json:"price_per_day"`
    Capacity    int     `json:"capacity"`
    Clients     []Client `gorm:"foreignKey:RoomID" json:"-"`
}

func (Room) TableName() string {
    return "room"
}

func (r *Room) PricePerPlace() float64 {
    return r.PricePerDay / float64(r.Capacity)
}

func (r *Room) IsSingle() bool {
    return r.RoomType == "single"
}
