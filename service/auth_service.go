package service

import (
	"final-project-2/entity"
	"final-project-2/repository/user_repository"
	"fmt"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Authentication() gin.HandlerFunc
}

type authService struct {
	userRepo user_repository.UserRepository
}

func NewAuthService(userRepo user_repository.UserRepository) AuthService {
	return &authService{userRepo}
}

func (a *authService) Authentication() gin.HandlerFunc {
	fmt.Println("authService.Authentication()")
	return func(ctx *gin.Context) {
		bearerToken := ctx.GetHeader("Authorization")

		var user entity.User

		if err := user.ValidateToken(bearerToken); err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		// fmt.Printf("%+v \n", user)

		result, err := a.userRepo.GetUserById(user.Id)
		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}
		
		ctx.Set("userData", result)
		ctx.Next()
	}
}
