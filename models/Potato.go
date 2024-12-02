package models

import "gorm.io/gorm"

type Potato struct {
	gorm.Model
	Img   string `json:"img"`
	Price uint   `json:"price"`
	Title string `json:"title"`
	Types []Type `json:"types" gorm:"many2many:potato_types;"`
	Sizes []Size `json:"sizes" gorm:"many2many:potato_sizes;"`
}
