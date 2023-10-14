package models

import (
	"gorm.io/gorm"

	"gorm.io/datatypes"
)

type User struct {
	gorm.Model
	Name      string         `json:"name"`
	Email     string         `gorm:"unique" json:"email"`
	Password  string         `json:"password"`
	Role      string         `json:"role"`
	Birthday  datatypes.Date `json:"birthday"`
	Entryday  datatypes.Date `json:"entryday"`
	Histories []History
}
