package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Potato struct {
	ID         uuid.UUID  `json:"id" gorm:"type:uuid;primaryKey"`
	Img        string     `json:"img"`
	Price      uint       `json:"price"`
	Title      string     `json:"title"`
	Rating     float32    `json:"rating"`
	Categories []Category `json:"category" gorm:"many2many:potato_categoris;"`

	CartPosition []CartPosition
	Types        []Type `json:"types" gorm:"many2many:potato_types;"`
	Sizes        []Size `json:"sizes" gorm:"many2many:potato_sizes;"`
}

func (p *Potato) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return
}
