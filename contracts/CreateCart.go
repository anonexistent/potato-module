package contracts

import (
	"potato-module/models"
)

type CreateCart struct {
	Positions []models.CartPosition `json:"positions"`
}
