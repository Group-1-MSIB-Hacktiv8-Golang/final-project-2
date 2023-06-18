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

type socialMediaHandler struct {
	socialMediaService service.SocialMediaService
}

func NewSocialMediaHandler(socialMediaService service.SocialMediaService) socialMediaHandler {
	return socialMediaHandler{
		socialMediaService: socialMediaService,
	}
}

func (s *socialMediaHandler) CreateNewSocialMedia(ctx *gin.Context) {
	var newSocialMediaRequest dto.NewSocialMediaRequest
	
	userData, ok := ctx.MustGet("userData").(*entity.User)
	if !ok {
		newError := errs.NewBadRequest("Failed to get user data")
		ctx.JSON(newError.Status(), newError)
		return
	}

	fmt.Println("Userdata di handler :", userData)

	if err := ctx.ShouldBindJSON(&newSocialMediaRequest); err != nil {
		errBindJSON := errs.NewUnprocessibleEntityError("invalid request body")
	
		ctx.JSON(errBindJSON.Status(), errBindJSON)
		return
	}

	//masukan id user ke newSocialMediaRequest
	newSocialMediaRequest.UserId = userData.Id
	
	result, err := s.socialMediaService.CreateNewSocialMedia(newSocialMediaRequest)
	
	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}
	
	ctx.JSON(result.StatusCode, result)
}

//get all social media by user id
func (s *socialMediaHandler) GetAllSocialMediaByUserId(ctx *gin.Context) {
	userData, ok := ctx.MustGet("userData").(*entity.User)
	if !ok {
		newError := errs.NewBadRequest("Failed to get user data")
		ctx.JSON(newError.Status(), newError)
		return
	}

	socialMedia, err := s.socialMediaService.GetAllSocialMediaByUserId(userData.Id)

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(socialMedia.StatusCode, socialMedia)
}


func (p socialMediaHandler) UpdateSocialMediaById(c *gin.Context) {
	var socialMediaRequest dto.UpdateSocialMediaRequest

	if err := c.ShouldBindJSON(&socialMediaRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		c.JSON(errBindJson.Status(), errBindJson)
		return
	}

	socialMediaId, err := helpers.GetParamId(c, "socialMediaId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	response, err := p.socialMediaService.UpdateSocialMediaById(socialMediaId, socialMediaRequest)
	
	fmt.Println("response di handler :", response, "err di handler :", err)

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(response.StatusCode, response)
}

func (p socialMediaHandler) DeleteSocialMediaById(c *gin.Context) {
	socialMediaId, err := helpers.GetParamId(c, "socialMediaId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	response, err := p.socialMediaService.DeleteSocialMediaById(socialMediaId)

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(response.StatusCode, response)
}
