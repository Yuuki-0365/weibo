package model

type Like struct {
	UserId  string `json:"user_id" gorm:"primaryKey"`
	NoteId  string `json:"note_id"`
	IsLiked bool   `json:"is_liked"`
}
