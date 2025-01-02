package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CartPosition struct {
	ID uuid.UUID `json:"id,omitempty" gorm:"type:uuid;primaryKey"`

	Potatos Potato `json:"potatoe"`
	Type    Type   `json:"type"`
	Size    Size   `json:"size"`
}

func (p *CartPosition) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return
}
