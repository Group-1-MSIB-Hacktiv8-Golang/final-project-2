package socialmedia_repository

import (
	"final-project-2/entity"
	"final-project-2/pkg/errs"
)

type SocialMediaRepository interface {
	CreateNewSocialMedia(socialMedia entity.SocialMedia) errs.MessageErr
	GetAllSocialMediaByUserId(userId int) ([]entity.SocialMedia, errs.MessageErr)
	UpdateSocialMediaById(payload entity.SocialMedia) errs.MessageErr
	DeleteSocialMediaById(socialMediaId int) errs.MessageErr
}
