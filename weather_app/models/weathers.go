package models

import (
	"gorm.io/gorm"
)

type Weather struct {
	gorm.Model
	Date        string  `json:"date"`
	Time        string  `json:"time"`
	City        string  `json:"city"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
}
