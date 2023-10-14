package models

import "gorm.io/gorm"

type Month struct {
	gorm.Model
	Month     int `json:"month"`
	Histories []History
}
