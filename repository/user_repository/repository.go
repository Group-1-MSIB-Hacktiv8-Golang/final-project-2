package user_repository

import (
	"final-project-2/entity"
	"final-project-2/pkg/errs"
)

type UserRepository interface {
	CreateNewUser(user entity.User) errs.MessageErr
	GetUserById(userId int) (*entity.User, errs.MessageErr)
	GetUserByEmail(userEmail string) (*entity.User, errs.MessageErr)
}
