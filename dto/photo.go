package dto

type NewPhotoRequest struct {
	UserId   int    `json:"user_id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
}

type NewPhotoResponse struct {
	Result     string `json:"result"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       struct {
		UserId   int    `json:"user_id"`
		Title    string `json:"title"`
		Caption  string `json:"caption"`
		PhotoUrl string `json:"photo_url"`
	} `json:"data"`
}

type NewPhotoResponseData struct {
	UserId   int    `json:"user_id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
}

type GetAllPhotoResponse struct {
	Result     string                    `json:"result"`
	StatusCode int                       `json:"status_code"`
	Message    string                    `json:"message"`
	Data       []GetAllPhotoResponseData `json:"data"`
}

type GetAllPhotoResponseData struct {
	Id        int    `json:"id"`
	UserId    int    `json:"user_id"`
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	PhotoUrl  string `json:"photo_url"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UpdatePhotoRequest struct {
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
}

type UpdatePhotoResponse struct {
	Result     string `json:"result"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       struct {
		Title    string `json:"title"`
		Caption  string `json:"caption"`
		PhotoUrl string `json:"photo_url"`
	} `json:"data"`
}

type UpdatePhotoResponseData struct {
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
}

type DeletePhotoResponse struct {
	Result     string `json:"result"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}
