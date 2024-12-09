package contracts

import "github.com/google/uuid"

// CreatePotatoBody represents the input structure for creating a potato
type CreatePotatoBody struct {
	Img   string      `json:"img"`
	Price uint        `json:"price"`
	Title string      `json:"title"`
	Types []uuid.UUID `json:"types"`
	Sizes []uuid.UUID `json:"sizes"`
}
