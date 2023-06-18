package photo_repository

import (
	"final-project-2/entity"
	"final-project-2/pkg/errs"
)

type PhotoRepository interface {
	CreateNewPhoto(photo entity.Photo) errs.MessageErr
	GetAllPhotoByUserId(userId int) ([]entity.Photo, errs.MessageErr)
	UpdatePhotoById(payload entity.Photo) errs.MessageErr
	DeletePhotoById(photoId int) errs.MessageErr
}
