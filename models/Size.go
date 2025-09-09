package models

type Size struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`

	Potatos []Potato `json:"potatoes,omitempty" gorm:"many2many:potato_sizes;"`
}
