package model

type LikeCount struct {
	UserId string `json:"user_id"`
	Count  int    `json:"count"`
}
