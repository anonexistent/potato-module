package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Type struct {
	ID           uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey"`
	Name         string         `json:"name"`
	CartPosition []CartPosition `json:"cart_positions" gorm:"many2many:cart_position_types;"`
	Potatos      []Potato       `json:"potatoes,omitempty" gorm:"many2many:potato_types;"`
}

func (p *Type) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return
}
