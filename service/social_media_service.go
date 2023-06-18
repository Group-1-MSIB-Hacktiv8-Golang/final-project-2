package service

import (
	"final-project-2/dto"
	"final-project-2/entity"
	"final-project-2/pkg/errs"
	"final-project-2/pkg/helpers"
	"final-project-2/repository/socialmedia_repository"
	"fmt"
	"net/http"
)

type SocialMediaService interface {
	CreateNewSocialMedia(payload dto.NewSocialMediaRequest) (*dto.NewSocialMediaResponse, errs.MessageErr)
	GetAllSocialMediaByUserId(userId int) (*dto.GetAllSocialMediaResponse, errs.MessageErr)
	UpdateSocialMediaById(socialMediaId int, socialMediaRequest dto.UpdateSocialMediaRequest) (*dto.UpdateSocialMediaResponse, errs.MessageErr)
	DeleteSocialMediaById(socialMediaId int) (*dto.DeleteSocialMediaResponse, errs.MessageErr)
}

type socialMediaService struct {
	socialMediaRepo socialmedia_repository.SocialMediaRepository
}

func NewSocialmediaService(socialMediaRepo socialmedia_repository.SocialMediaRepository) SocialMediaService {
	return &socialMediaService{
		socialMediaRepo: socialMediaRepo,
	}
}

//CreateNewSocialMedia
func (s *socialMediaService) CreateNewSocialMedia(payload dto.NewSocialMediaRequest) (*dto.NewSocialMediaResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(payload)

	if err != nil {
		return nil, err
	}

	socialMedia := entity.SocialMedia{
		UserId:   		payload.UserId,
		Name:   		payload.Name,
		SocialMediaUrl: payload.SocialMediaUrl,
	}

	err = s.socialMediaRepo.CreateNewSocialMedia(socialMedia)


	if err != nil {
		return nil, err
	}

	response := dto.NewSocialMediaResponse{
		Result:     "success",
		StatusCode: 200,
		Message:    "successfully created new social media",
		Data: dto.NewSocialMediaResponseData{
			UserId:   		socialMedia.UserId,
			Name:   		socialMedia.Name,
			SocialMediaUrl: socialMedia.SocialMediaUrl,
		},
	}

	return &response, nil
}

func (p *socialMediaService) GetAllSocialMediaByUserId(userId int) (*dto.GetAllSocialMediaResponse, errs.MessageErr) {
	//execute query
	socialMedias, err := p.socialMediaRepo.GetAllSocialMediaByUserId(userId)

	fmt.Println("socialMedias", socialMedias)
	if err != nil {
		return nil, err
	}

	//get all data
	var socialMediaResponses []dto.GetAllSocialMediaResponseData
	for _, socialMedia := range socialMedias {
		socialMediaResponse := dto.GetAllSocialMediaResponseData{
			Id: 			socialMedia.Id,
			UserId:   		socialMedia.UserId,
			Name:   		socialMedia.Name,
			SocialMediaUrl: socialMedia.SocialMediaUrl,
		}

		socialMediaResponses = append(socialMediaResponses, socialMediaResponse)
	}

	response := dto.GetAllSocialMediaResponse{
		Result:     "success",
		StatusCode: 200,
		Message:    "successfully get all social media",
		Data:       socialMediaResponses,
	}

	return &response, nil
}

func (p *socialMediaService) UpdateSocialMediaById(socialMediaId int, socialMediaRequest dto.UpdateSocialMediaRequest) (*dto.UpdateSocialMediaResponse, errs.MessageErr) {

	err := helpers.ValidateStruct(socialMediaRequest)

	if err != nil {
		return nil, err
	}

	payload := entity.SocialMedia{
		Id: 			socialMediaId,
		Name:   		socialMediaRequest.Name,
		SocialMediaUrl: socialMediaRequest.SocialMediaUrl,
	}

	err = p.socialMediaRepo.UpdateSocialMediaById(payload)

	fmt.Println("err di service", err)

	if err != nil {
		return nil, err
	}

	response := dto.UpdateSocialMediaResponse{
		Result:     "success",
		StatusCode: http.StatusOK,
		Message:    "successfully update socialMedia",
		Data: dto.UpdateSocialMediaResponseData{
			Name:   		payload.Name,
			SocialMediaUrl: payload.SocialMediaUrl,
		},
	}

	return &response, nil
}

//delete socialMedia
func (p *socialMediaService) DeleteSocialMediaById(socialMediaId int) (*dto.DeleteSocialMediaResponse, errs.MessageErr) {
	err := p.socialMediaRepo.DeleteSocialMediaById(socialMediaId)

	if err != nil {
		return nil, err
	}

	response := dto.DeleteSocialMediaResponse{
		Result:     "success",
		StatusCode: http.StatusOK,
		Message:    "successfully delete socialMedia",
	}

	return &response, nil
}