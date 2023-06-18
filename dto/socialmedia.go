package dto

type NewSocialMediaRequest struct {
	UserId         int    `json:"user_id"`
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
}

type NewSocialMediaResponse struct {
	Result     string `json:"result"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       struct {
		UserId         int    `json:"user_id"`
		Name           string `json:"name"`
		SocialMediaUrl string `json:"social_media_url"`
	} `json:"data"`
}

type NewSocialMediaResponseData struct {
	UserId         int    `json:"user_id"`
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
}

type GetAllSocialMediaResponse struct {
	Result     string                          `json:"result"`
	StatusCode int                             `json:"status_code"`
	Message    string                          `json:"message"`
	Data       []GetAllSocialMediaResponseData `json:"data"`
}

type GetAllSocialMediaResponseData struct {
	Id             int    `json:"id"`
	UserId         int    `json:"user_id"`
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

// UpdateSocialMediaRequest
type UpdateSocialMediaRequest struct {
	UserId         int    `json:"user_id"`
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
}

// UpdateSocialMediaResponse
type UpdateSocialMediaResponse struct {
	Result     string `json:"result"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       struct {
		UserId         int    `json:"user_id"`
		Name           string `json:"name"`
		SocialMediaUrl string `json:"social_media_url"`
	} `json:"data"`
}

// UpdateSocialMediaResponseData
type UpdateSocialMediaResponseData struct {
	UserId         int    `json:"user_id"`
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
}

type DeleteSocialMediaResponse struct {
	Result     string `json:"result"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}
