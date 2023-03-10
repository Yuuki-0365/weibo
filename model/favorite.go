package model

type Favorite struct {
	CreateAt   string `json:"create_at"`
	UpdateAt   string `json:"update_at"`
	UserId     string `json:"user_id"`
	NoteId     uint   `json:"note_id"`
	IsFavorite bool   `json:"is_favorite"`
}
