package models

type Cart struct {
	ID        uint            `json:"id" gorm:"primaryKey"`
	Positions []*CartPosition `json:"positions"`
}
