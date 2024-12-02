package models

import "gorm.io/gorm"

type Type struct {
	gorm.Model
	Name    string    `json:"name"`
	Potatos []*Potato `gorm:"many2many:potato_types;"`
}
