package contracts

import "github.com/google/uuid"

type PositionIdBody struct {
	ID uuid.UUID `json:"id"`
}
