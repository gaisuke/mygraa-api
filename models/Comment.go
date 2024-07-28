package models

type Comment struct {
	ID        uint   `json:"id"`
	UserId    uint   `json:"user_id"`
	PhotoId   uint   `json:"photo_id"`
	Message   string `json:"message"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
