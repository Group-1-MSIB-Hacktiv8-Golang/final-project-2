package handler

import (
	"final-project-2/dto"
	"final-project-2/entity"
	"final-project-2/pkg/errs"
	"final-project-2/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

type commentHandler struct {
	commentService service.CommentService
}

func NewCommentHandler(commentService service.CommentService) commentHandler {
	return commentHandler{
		commentService: commentService,
	}
}

func (c *commentHandler) CreateNewComment(ctx *gin.Context) {
	var newCommentRequest dto.NewCommentRequest
	
	userData, ok := ctx.MustGet("userData").(*entity.User)
	if !ok {
		newError := errs.NewBadRequest("Failed to get user data")
		ctx.JSON(newError.Status(), newError)
		return
	}

	fmt.Println("Userdata di handler :", userData)

	if err := ctx.ShouldBindJSON(&newCommentRequest); err != nil {
		errBindJSON := errs.NewUnprocessibleEntityError("invalid request body")
	
		ctx.JSON(errBindJSON.Status(), errBindJSON)
		return
	}

	//masukan id user ke newPhotoRequest
	newCommentRequest.UserId = userData.Id
	
	result, err := c.commentService.CreateNewComment(newCommentRequest)
	
	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}
	
	ctx.JSON(result.StatusCode, result)
}