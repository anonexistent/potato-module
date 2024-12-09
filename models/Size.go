package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Size struct {
	ID      uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Name    string    `json:"name"`
	Potatos []Potato  `json:"potatoes,omitempty" gorm:"many2many:potato_sizes;"`
}

func (p *Size) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return
}
