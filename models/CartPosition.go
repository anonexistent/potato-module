package models

type CartPosition struct {
	ID uint `gorm:"primaryKey"`

	CartId uint  `json:"cartId"`
	Cart   *Cart `json:"omitempty"`

	PotatoId uint    `json:"potatoId"`
	Potato   *Potato `json:"omitempty"`

	SizeId uint  `json:"sizeId"`
	Size   *Size `json:"omitempty"`

	TypeId uint  `json:"typeId"`
	Type   *Type `json:"omitempty"`
}
