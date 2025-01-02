package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Cart struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey"`
	Positions []CartPosition `json:"positions"`
}

func (p *Cart) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return
}
