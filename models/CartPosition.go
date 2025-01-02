package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CartPosition struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	CartId uuid.UUID `json:"cartId"`
	Cart   *Cart     `json:"omitempty"`

	PotatoId uuid.UUID `json:"potatoId"`
	Potato   *Potato   `json:"omitempty"`

	SizeId uuid.UUID `json:"sizeId"`
	Size   *Size     `json:"omitempty"`

	TypeId uuid.UUID `json:"typeId"`
	Type   *Type     `json:"omitempty"`
}

func (p *CartPosition) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return
}
