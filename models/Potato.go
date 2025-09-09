package models

type Potato struct {
	ID         uint       `json:"id" gorm:"primaryKey"`
	Img        string     `json:"img"`
	Price      uint       `json:"price"`
	Title      string     `json:"title"`
	Rating     float32    `json:"rating"`
	Categories []Category `json:"category" gorm:"many2many:potato_categoris;"`

	Types []Type `json:"types" gorm:"many2many:potato_types;"`
	Sizes []Size `json:"sizes" gorm:"many2many:potato_sizes;"`
}
