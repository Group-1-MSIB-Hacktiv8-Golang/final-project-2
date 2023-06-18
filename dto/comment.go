package dto

type NewCommentRequest struct {
	UserId  int    `json:"user_id"`
	PhotoId int    `json:"photo_id"`
	Message string `json:"message"`
}

type NewCommentResponse struct {
	Result     string `json:"result"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       struct {
		UserId  int    `json:"user_id"`
		PhotoId int    `json:"photo_id"`
		Message string `json:"message"`
	} `json:"data"`
}

type NewCommentResponseData struct {
	UserId  int    `json:"user_id"`
	PhotoId int    `json:"photo_id"`
	Message string `json:"message"`
}