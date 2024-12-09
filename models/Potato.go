package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Potato struct {
	ID    uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Img   string    `json:"img"`
	Price uint      `json:"price"`
	Title string    `json:"title"`
	Types []Type    `json:"types" gorm:"many2many:potato_types;"`
	Sizes []Size    `json:"sizes" gorm:"many2many:potato_sizes;"`
}

// BeforeCreate будет вызываться перед созданием записи
func (p *Potato) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New() // Генерируем новый UUID
	return
}
