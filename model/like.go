package model

type Like struct {
	UserId  string `json:"user_id"`
	NoteId  uint   `json:"note_id"`
	IsLiked bool   `json:"is_liked"`
}
