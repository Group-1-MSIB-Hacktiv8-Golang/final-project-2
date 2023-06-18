package comment_repository

import (
	"final-project-2/entity"
	"final-project-2/pkg/errs"
)

type CommentRepository interface {
	CreateNewComment(comment entity.Comment) errs.MessageErr
	GetAllCommentByUserId(userId int) ([]entity.Comment, errs.MessageErr)
	UpdateCommentById(commentId int, comment entity.Comment) errs.MessageErr
	DeleteCommentById(commentId int) errs.MessageErr
}
