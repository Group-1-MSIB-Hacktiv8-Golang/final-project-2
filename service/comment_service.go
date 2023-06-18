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