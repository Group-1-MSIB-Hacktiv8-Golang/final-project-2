package service

import (
	"final-project-2/dto"
	"final-project-2/entity"
	"final-project-2/pkg/errs"
	"final-project-2/pkg/helpers"
	"final-project-2/repository/user_repository"
	"net/http"
)

type UserService interface {
	CreateNewUser(payload dto.NewUserRequest) (*dto.NewUserResponse, errs.MessageErr)
	Login(loginUserRequest dto.LoginUserRequest) (*dto.LoginResponse, errs.MessageErr)
	UpdateUser(user *entity.User, payload *dto.UpdateUserRequest) (*dto.UpdateUserResponse, errs.MessageErr)
	DeleteUser(user *entity.User) (*dto.DeleteUserResponse, errs.MessageErr)
}

type userService struct {
	userRepo user_repository.UserRepository
}

func NewUserService(userRepo user_repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (u *userService) Login(loginUserRequest dto.LoginUserRequest) (*dto.LoginResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(loginUserRequest)

	if err != nil {
		return nil, err
	}

	user, err := u.userRepo.GetUserByEmail(loginUserRequest.Email)

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, errs.NewBadRequest("invalid email/password")
		}
		return nil, err
	}

	isValidPassword := user.ComparePassword(loginUserRequest.Password)

	if !isValidPassword {
		return nil, errs.NewBadRequest("invalid email/password")
	}

	token := user.GenerateToken()

	response := dto.LoginResponse{
		Result:     "success",
		StatusCode: http.StatusOK,
		Message:    "successfully logged in",
		Data: dto.TokenResponse{
			Token: token,
		},
	}

	return &response, nil
}

func (u *userService) CreateNewUser(payload dto.NewUserRequest) (*dto.NewUserResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(payload)

	if err != nil {
		return nil, err
	}

	user := entity.User{
		Email:    payload.Email,
		Password: payload.Password,
		Username: payload.Username,
		Age:      payload.Age,
	}

	err = user.HashPassword()

	if err != nil {
		return nil, err
	}

	err = u.userRepo.CreateNewUser(user)


	if err != nil {
		return nil, err
	}

	response := dto.NewUserResponse{
		Result:     "success",
		StatusCode: http.StatusCreated,
		Message:    "user registered successfully",
	}

	return &response, nil
}


//UpdateUser
func (u *userService) UpdateUser(user *entity.User, payload *dto.UpdateUserRequest) (*dto.UpdateUserResponse, errs.MessageErr) {

	newUser := payload.ToEntity()

	updatedUser, err := u.userRepo.UpdateUser(newUser, user.Id)

	if err != nil {
		return nil, err
	}

	response := &dto.UpdateUserResponse{
		Result:     "success",
		StatusCode: http.StatusOK,
		Message:    "user updated successfully",
		Data: dto.UpdateUserResponseData{
			Id:        updatedUser.Id,
			Email:     updatedUser.Email,
			Username:  updatedUser.Username,
			Age:       updatedUser.Age,
		},
	}

	return response, nil
}

func (u *userService) DeleteUser(user *entity.User) (*dto.DeleteUserResponse, errs.MessageErr) {
	err := u.userRepo.DeleteUser(user.Id)

	if err != nil {
		return nil, err
	}

	response := &dto.DeleteUserResponse{
		Result:     "success",
		StatusCode: http.StatusOK,
		Message:    "user deleted successfully",
	}

	return response, nil
}