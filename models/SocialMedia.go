package models

type SocialMedia struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
	UserId         uint   `json:"user_id"`
}
