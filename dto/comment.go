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

// GetAllCommentResponse
type GetAllCommentResponse struct {
	Result     string                      `json:"result"`
	StatusCode int                         `json:"status_code"`
	Message    string                      `json:"message"`
	Data       []GetAllCommentResponseData `json:"data"`
}

type GetAllCommentResponseData struct {
	Id        int    `json:"id"`
	UserId    int    `json:"user_id"`
	PhotoId   int    `json:"photo_id"`
	Message   string `json:"message"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UpdateCommentRequest struct {
	UserId  int    `json:"user_id"`
	Message string `json:"message"`
	PhotoId int    `json:"photo_id"`
}

type UpdateCommentResponse struct {
	Result     string `json:"result"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       struct {
		Message string `json:"message"`
	} `json:"data"`
}

type UpdateCommentResponseData struct {
	Message string `json:"message"`
}

type DeleteCommentResponse struct {
	Result     string `json:"result"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}