package service

import (
	"final-project-2/dto"
	"final-project-2/entity"
	"final-project-2/pkg/errs"
	"final-project-2/pkg/helpers"
	"final-project-2/repository/comment_repository"
)

type CommentService interface {
	CreateNewComment(payload dto.NewCommentRequest) (*dto.NewCommentResponse, errs.MessageErr)
	GetAllCommentByUserId(userId int) (*dto.GetAllCommentResponse, errs.MessageErr)
	UpdateCommentById(commentId int, commentRequest dto.UpdateCommentRequest) (*dto.UpdateCommentResponse, errs.MessageErr)


}

type commentService struct {
	commentRepo comment_repository.CommentRepository
}

func NewCommentService(commentRepo comment_repository.CommentRepository) CommentService {
	return &commentService{
		commentRepo: commentRepo,
	}
}

func (c *commentService) CreateNewComment(payload dto.NewCommentRequest) (*dto.NewCommentResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(payload)

	if err != nil {
		return nil, err
	}

	comment := entity.Comment{
		UserId:   payload.UserId,
		PhotoId:    payload.PhotoId,
		Message: payload.Message,
	}

	err = c.commentRepo.CreateNewComment(comment)


	if err != nil {
		return nil, err
	}

	response := dto.NewCommentResponse{
		Result:     "success",
		StatusCode: 200,
		Message:    "successfully created new comment",
		Data: dto.NewCommentResponseData{
			UserId:   comment.UserId,
			PhotoId:    comment.PhotoId,
			Message: comment.Message,
		},
	}

	return &response, nil
}

func (c *commentService) GetAllCommentByUserId(userId int) (*dto.GetAllCommentResponse, errs.MessageErr) {
	comments, err := c.commentRepo.GetAllCommentByUserId(userId)

	if err != nil {
		return nil, err
	}

	var commentsResponse []dto.GetAllCommentResponseData

	for _, comment := range comments {
		commentsResponse = append(commentsResponse, dto.GetAllCommentResponseData{
			Id:        comment.Id,
			UserId:    comment.UserId,
			PhotoId:   comment.PhotoId,
			Message:   comment.Message,
		})
	}

	response := dto.GetAllCommentResponse{
		Result:     "success",
		StatusCode: 200,
		Message:    "successfully get all comment",
		Data:       commentsResponse,
	}

	return &response, nil
}

func (c *commentService) UpdateCommentById(commentId int, commentRequest dto.UpdateCommentRequest) (*dto.UpdateCommentResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(commentRequest)

	if err != nil {
		return nil, err
	}

	comment := entity.Comment{
		Id:       commentId,
		UserId:   commentRequest.UserId,
		PhotoId:   commentRequest.PhotoId,
		Message: commentRequest.Message,
	}

	err = c.commentRepo.UpdateCommentById(commentId, comment)

	if err != nil {
		return nil, err
	}

	response := dto.UpdateCommentResponse{
		Result:     "success",
		StatusCode: 200,
		Message:    "successfully update comment",
		Data: dto.UpdateCommentResponseData{
			Message: comment.Message,
		},
	}

	return &response, nil
}