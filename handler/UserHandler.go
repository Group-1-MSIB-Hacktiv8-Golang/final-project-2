package handler

import (
	"final-project-2/dto"
	"final-project-2/entity"
	"final-project-2/pkg/errs"
	"final-project-2/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) userHandler {
	return userHandler{
		userService: userService,
	}
}

func (uh *userHandler) Register(ctx *gin.Context) {
	var newUserRequest dto.NewUserRequest
	
	if err := ctx.ShouldBindJSON(&newUserRequest); err != nil {
		errBindJSON := errs.NewUnprocessibleEntityError("invalid request body")
		
		ctx.JSON(errBindJSON.Status(), errBindJSON)
		return
	}

	result, err := uh.userService.CreateNewUser(newUserRequest)

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(result.StatusCode, result)
}

func (uh *userHandler) Login(ctx *gin.Context) {
	var loginUserRequest dto.LoginUserRequest

	if err := ctx.ShouldBindJSON(&loginUserRequest); err != nil {
		errBindJSON := errs.NewUnprocessibleEntityError("invalid request body")

		ctx.JSON(errBindJSON.Status(), errBindJSON)
		return
	}

	result, err := uh.userService.Login(loginUserRequest)

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(result.StatusCode, result)
}

func (uh *userHandler) UpdateUser(ctx *gin.Context) {
	var updateUserRequest dto.UpdateUserRequest

	userData, ok := ctx.MustGet("userData").(*entity.User)
	if !ok {
		newError := errs.NewBadRequest("Failed to get user data")
		ctx.JSON(newError.Status(), newError)
		return
	}

	fmt.Println("user data", userData)

	if err := ctx.ShouldBindJSON(&updateUserRequest); err != nil {
		errBindJSON := errs.NewUnprocessibleEntityError("invalid request body")

		ctx.JSON(errBindJSON.Status(), errBindJSON)
		return
	}

	fmt.Println("update user request", updateUserRequest)

	result, err := uh.userService.UpdateUser(userData, &updateUserRequest)
	fmt.Println("result", result)
	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(result.StatusCode, result)
}

//delete user
func (uh *userHandler) DeleteUser(ctx *gin.Context) {
	userData, ok := ctx.MustGet("userData").(*entity.User)
	if !ok {
		newError := errs.NewBadRequest("Failed to get user data")
		ctx.JSON(newError.Status(), newError)
		return
	}

	result, err := uh.userService.DeleteUser(userData)
	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	// Mengirim respons JSON dengan kode status dan hasil yang diterima dari userService.
	ctx.JSON(result.StatusCode, result)

}