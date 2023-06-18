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

func (c *commentHandler) GetAllCommentByUserId(ctx *gin.Context) {
	userData, ok := ctx.MustGet("userData").(*entity.User)
	if !ok {
		newError := errs.NewBadRequest("Failed to get user data")
		ctx.JSON(newError.Status(), newError)
		return
	}

	fmt.Println("Userdata di handler :", userData)

	result, err := c.commentService.GetAllCommentByUserId(userData.Id)
	
	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}
	
	ctx.JSON(result.StatusCode, result)
}

//UpdateCommentById
func (c *commentHandler) UpdateCommentById(ctx *gin.Context) {
	var updateCommentRequest dto.UpdateCommentRequest
	
	userData, ok := ctx.MustGet("userData").(*entity.User)
	if !ok {
		newError := errs.NewBadRequest("Failed to get user data")
		ctx.JSON(newError.Status(), newError)
		return
	}

	fmt.Println("Userdata di handler :", userData)

	if err := ctx.ShouldBindJSON(&updateCommentRequest); err != nil {
		errBindJSON := errs.NewUnprocessibleEntityError("invalid request body")
	
		ctx.JSON(errBindJSON.Status(), errBindJSON)
		return
	}

	//masukan id user ke newPhotoRequest
	commentId, err := helpers.GetParamId(ctx, "commentId")
	
	result, err := c.commentService.UpdateCommentById(commentId, updateCommentRequest)
	
	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}
	
	ctx.JSON(result.StatusCode, result)
}