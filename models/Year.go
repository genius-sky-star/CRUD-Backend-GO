package models

import "gorm.io/gorm"

type Year struct {
	gorm.Model
	Year      int `json:"year"`
	Histories []History
}
