package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	ID      uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Name    string    `json:"title"`
	Potatos []Potato  `json:"potatoes,omitempty" gorm:"many2many:potato_categoris;"`
}

func (p *Category) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return
}
