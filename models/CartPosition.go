package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CartPosition struct {
	ID uuid.UUID `json:"id,omitempty" gorm:"type:uuid;primaryKey"`

	Potatos []Potato  `json:"potatoes" gorm:"many2many:cart_position_potatos;"`
	TypeID  uuid.UUID `json:"type_id" gorm:"type:uuid"`      // Добавьте поле для внешнего ключа
	Type    Type      `json:"type" gorm:"foreignKey:TypeID"` // Укажите внешний ключ
	SizeID  uuid.UUID `json:"size_id" gorm:"type:uuid"`      // Если у вас есть связь с Size
	Size    Size      `json:"size" gorm:"foreignKey:SizeID"` // Укажите внешний ключ
}

func (p *CartPosition) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return
}
