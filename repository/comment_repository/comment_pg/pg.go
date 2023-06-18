package comment_pg

import (
	"database/sql"
	"final-project-2/entity"
	"final-project-2/pkg/errs"
	"final-project-2/repository/comment_repository"
)

const (
	createNewComment = `
		INSERT INTO "comments"
		(
			user_id,
			photo_id,
			message
		)
		VALUES ($1, $2, $3);
	`
)
type commentPG struct {
	db *sql.DB
}

func NewCommentPG(db *sql.DB) comment_repository.CommentRepository {
	return &commentPG{
		db: db,
	}
}

func (c *commentPG) CreateNewComment(comment entity.Comment) errs.MessageErr {
	_, err := c.db.Exec(createNewComment, comment.UserId, comment.PhotoId, comment.Message)
	if err != nil {
		return errs.NewInternalServerError("Database Error")
	}

	return nil
}