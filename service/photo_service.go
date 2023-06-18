package service

import (
	"final-project-2/dto"
	"final-project-2/entity"
	"final-project-2/pkg/errs"
	"final-project-2/pkg/helpers"
	"final-project-2/repository/photo_repository"
	"fmt"
	"net/http"
)

type PhotoService interface {
	CreateNewPhoto(payload dto.NewPhotoRequest) (*dto.NewPhotoResponse, errs.MessageErr)
	GetAllPhotoByUserId(userId int) (*dto.GetAllPhotoResponse, errs.MessageErr)
	UpdatePhotoById(photoId int, photoRequest dto.UpdatePhotoRequest) (*dto.UpdatePhotoResponse, errs.MessageErr)
	DeletePhotoById(photoId int) (*dto.DeletePhotoResponse, errs.MessageErr)
}

type photoService struct {
	photoRepo photo_repository.PhotoRepository
}

func NewPhotoService(photoRepo photo_repository.PhotoRepository) PhotoService {
	return &photoService{
		photoRepo: photoRepo,
	}
}

func (p *photoService) CreateNewPhoto(payload dto.NewPhotoRequest) (*dto.NewPhotoResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(payload)

	if err != nil {
		return nil, err
	}

	photo := entity.Photo{
		UserId:   payload.UserId,
		Title:    payload.Title,
		Caption:  payload.Caption,
		PhotoUrl: payload.PhotoUrl,
	}

	err = p.photoRepo.CreateNewPhoto(photo)


	if err != nil {
		return nil, err
	}

	response := dto.NewPhotoResponse{
		Result:     "success",
		StatusCode: 200,
		Message:    "successfully created new photo",
		Data: dto.NewPhotoResponseData{
			UserId:   photo.UserId,
			Title:    photo.Title,
			Caption:  photo.Caption,
			PhotoUrl: photo.PhotoUrl,
		},
	}

	return &response, nil
}

//get all photo by user id
func (p *photoService) GetAllPhotoByUserId(userId int) (*dto.GetAllPhotoResponse, errs.MessageErr) {
	//execute query
	photos, err := p.photoRepo.GetAllPhotoByUserId(userId)

	fmt.Println("photos", photos)
	if err != nil {
		return nil, err
	}

	//get all data
	var photoResponses []dto.GetAllPhotoResponseData
	for _, photo := range photos {
		photoResponse := dto.GetAllPhotoResponseData{
			Id:        photo.Id,
			UserId:    photo.UserId,
			Title:     photo.Title,
			Caption:   photo.Caption,
			PhotoUrl:  photo.PhotoUrl,
		}

		photoResponses = append(photoResponses, photoResponse)
	}

	response := dto.GetAllPhotoResponse{
		Result:     "success",
		StatusCode: 200,
		Message:    "successfully get all photo",
		Data:       photoResponses,
	}

	return &response, nil
}

//update photo
func (p *photoService) UpdatePhotoById(photoId int, photoRequest dto.UpdatePhotoRequest) (*dto.UpdatePhotoResponse, errs.MessageErr) {

	err := helpers.ValidateStruct(photoRequest)

	if err != nil {
		return nil, err
	}

	payload := entity.Photo{
		Title:    photoRequest.Title,
		Caption:  photoRequest.Caption,
		PhotoUrl: photoRequest.PhotoUrl,
		Id:       photoId,
	}

	err = p.photoRepo.UpdatePhotoById(payload)

	fmt.Println("err di service", err)

	if err != nil {
		return nil, err
	}

	response := dto.UpdatePhotoResponse{
		Result:     "success",
		StatusCode: http.StatusOK,
		Message:    "successfully update photo",
		Data: dto.UpdatePhotoResponseData{
			Title:    payload.Title,
			Caption:  payload.Caption,
			PhotoUrl: payload.PhotoUrl,
		},
	}

	return &response, nil
}

//delete photo
func (p *photoService) DeletePhotoById(photoId int) (*dto.DeletePhotoResponse, errs.MessageErr) {
	err := p.photoRepo.DeletePhotoById(photoId)

	if err != nil {
		return nil, err
	}

	response := dto.DeletePhotoResponse{
		Result:     "success",
		StatusCode: http.StatusOK,
		Message:    "successfully delete photo",
	}

	return &response, nil
}