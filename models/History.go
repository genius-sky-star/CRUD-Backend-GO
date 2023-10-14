package models

import "gorm.io/gorm"

type History struct {
	gorm.Model
	Amount  int `json:"amount"`
	UserID  uint
	YearID  uint
	MonthID uint
}
