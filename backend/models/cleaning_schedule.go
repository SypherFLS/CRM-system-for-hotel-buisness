package models

import (
    "gorm.io/gorm"
    "time"
)

type CleaningSchedule struct {
    gorm.Model
    EmployeeID   uint
    RoomID       uint
    CleaningDate time.Time
    Employee     Employee `gorm:"foreignKey:EmployeeID"`
    Room         Room     `gorm:"foreignKey:RoomID"`
}
func (CleaningSchedule) TableName() string {
    return "cleaning_schedule" 
}