package dto

import "final-project-2/entity"

type NewUserRequest struct {
	// Email    string `json:"email" valid:"required~email cannot be empty"`
	//email cannot be empty and must be a valid email and unique
	Email    string `json:"email" valid:"required~email cannot be empty,email~email must be a valid email"`
	Password string `json:"password" valid:"required~password cannot be empty"`
	Username string `json:"username" valid:"required~username cannot be empty"`
	Age      int    `json:"age" valid:"required~age cannot be empty"`
}

type LoginUserRequest struct {
	Email    string `json:"email" valid:"required~email cannot be empty"`
	Password string `json:"password" valid:"required~password cannot be empty"`
}

type NewUserResponse struct {
	Result     string `json:"result"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type LoginResponse struct {
	Result     string        `json:"result"`
	StatusCode int           `json:"statusCode"`
	Message    string        `json:"message"`
	Data       TokenResponse `json:"data"`
}

//ToEntity

type UpdateUserRequest struct {
	Email    string `json:"email" valid:"required~email cannot be empty,email~email must be a valid email"`
	Username string `json:"username" valid:"required~username cannot be empty"`
}

func (n *UpdateUserRequest) ToEntity() *entity.User {
	return &entity.User{
		Email:    n.Email,
		Username: n.Username,
	}
}

type UpdateUserResponse struct {
	//id
	//email
	//username
	//age

	Result     string `json:"result"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`

	Data struct {
		Id        int    `json:"id"`
		Email     string `json:"email"`
		Username  string `json:"username"`
		Age       int    `json:"age"`
	} `json:"data"`
}

type UpdateUserResponseData struct {
	Id        int    `json:"id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Age       int    `json:"age"`
}

type DeleteUserResponse struct {
	Result     string `json:"result"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}