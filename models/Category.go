package models

type Category struct {
	ID      uint     `json:"id" gorm:"primaryKey"`
	Name    string   `json:"title"`
	Potatos []Potato `json:"potatoes,omitempty" gorm:"many2many:potato_categoris;"`
}
