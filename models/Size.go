package models

import "gorm.io/gorm"

type Size struct {
	gorm.Model
	Name    string    `json:"name"`
	Potatos []*Potato `gorm:"many2many:potato_sizes;"`
}
