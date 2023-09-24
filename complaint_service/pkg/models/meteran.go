package models

import "gorm.io/gorm"

type Meteran struct {
	gorm.Model
	User        User `json:"user"`
	UserID      uint
	MeteranCode string  `json:"meteran_code"`
	Address     string  `json:"address"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
}
