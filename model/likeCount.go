package model

type LikeCount struct {
	UserId string `json:"user_id" gorm:"primaryKey"`
	Count  int    `json:"count"`
}
