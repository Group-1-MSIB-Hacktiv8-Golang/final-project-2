package entity

import "time"

type SocialMedia struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	Name      string    `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}