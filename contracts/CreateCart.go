package contracts

import (
	"potato-module/models"
)

type CreateCart struct {
	Position models.CartPosition `json:"position"`
}
