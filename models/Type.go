package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Type struct {
	gorm.Model
	ID      uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Name    string    `json:"name"`
	Potatos []*Potato `gorm:"many2many:potato_types;"`
}

// BeforeCreate будет вызываться перед созданием записи
func (p *Type) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New() // Генерируем новый UUID
	return
}
