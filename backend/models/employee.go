package models

import "gorm.io/gorm"

type Employee struct {
    gorm.Model
    FullName string `json:"full_name"`
    Floors   string `json:"floors"`
    WeekDays string `json:"week_days"`
}
func (Employee) TableName() string {
    return "employee"
}