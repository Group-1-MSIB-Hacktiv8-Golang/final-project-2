package handler

import (
	"final-project-2/dto"
	"final-project-2/entity"
	"final-project-2/pkg/errs"
	"final-project-2/pkg/helpers"
	"final-project-2/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

type photoHandler struct {
	photoService service.PhotoService
}

func NewPhotoHandler(photoService service.PhotoService) photoHandler {
	return photoHandler{
		photoService: photoService,
	}
}

func (p *photoHandler) CreateNewPhoto(ctx *gin.Context) {
	var newPhotoRequest dto.NewPhotoRequest
	
	userData, ok := ctx.MustGet("userData").(*entity.User)
	if !ok {
		newError := errs.NewBadRequest("Failed to get user data")
		ctx.JSON(newError.Status(), newError)
		return
	}

	fmt.Println("Userdata di handler :", userData)

	if err := ctx.ShouldBindJSON(&newPhotoRequest); err != nil {
		errBindJSON := errs.NewUnprocessibleEntityError("invalid request body")
	
		ctx.JSON(errBindJSON.Status(), errBindJSON)
		return
	}

	//masukan id user ke newPhotoRequest
	newPhotoRequest.UserId = userData.Id

	fmt.Println("newPhotoRequest di handler :",newPhotoRequest)
	
	result, err := p.photoService.CreateNewPhoto(newPhotoRequest)

	fmt.Println("result di handler :",result , "err di handler :", err)
	
	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}
	
	ctx.JSON(result.StatusCode, result)
}

//get all photo by user id
func (p *photoHandler) GetAllPhotoByUserId(ctx *gin.Context) {
	userData, ok := ctx.MustGet("userData").(*entity.User)
	if !ok {
		newError := errs.NewBadRequest("Failed to get user data")
		ctx.JSON(newError.Status(), newError)
		return
	}

	photo, err := p.photoService.GetAllPhotoByUserId(userData.Id)

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(photo.StatusCode, photo)
}

//UpdatePhoto
func (p photoHandler) UpdatePhotoById(c *gin.Context) {
	var photoRequest dto.UpdatePhotoRequest

	if err := c.ShouldBindJSON(&photoRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		c.JSON(errBindJson.Status(), errBindJson)
		return
	}

	photoId, err := helpers.GetParamId(c, "photoId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	response, err := p.photoService.UpdatePhotoById(photoId, photoRequest)
	
	fmt.Println("response di handler :", response, "err di handler :", err)

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(response.StatusCode, response)
}

//DeletePhotoById
func (p photoHandler) DeletePhotoById(c *gin.Context) {
	photoId, err := helpers.GetParamId(c, "photoId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	response, err := p.photoService.DeletePhotoById(photoId)

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(response.StatusCode, response)
}