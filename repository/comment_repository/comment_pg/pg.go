package comment_pg

import (
	"database/sql"
	"final-project-2/entity"
	"final-project-2/pkg/errs"
	"final-project-2/repository/comment_repository"
	"fmt"
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
	fmt.Println(err)
	if err != nil {
		return errs.NewInternalServerError("Database Error")
	}

	return nil
}

func (c *commentPG) GetAllCommentByUserId(userId int) ([]entity.Comment, errs.MessageErr) {
	var comments []entity.Comment

	rows, err := c.db.Query("SELECT * FROM comments WHERE user_id = $1", userId)

	if err != nil {
		return nil, errs.NewInternalServerError("Database Error")
	}

	for rows.Next() {
		var comment entity.Comment
		err := rows.Scan(&comment.Id, &comment.UserId, &comment.PhotoId, &comment.Message, &comment.CreatedAt, &comment.UpdatedAt)

		if err != nil {
			return nil, errs.NewInternalServerError("Database Error")
		}

		comments = append(comments, comment)
	}

	return comments, nil
}

func (c *commentPG) UpdateCommentById(commentId int, comment entity.Comment) errs.MessageErr {
	_, err := c.db.Exec("UPDATE comments SET message = $1 WHERE id = $2", comment.Message, commentId)

	if err != nil {
		return errs.NewInternalServerError("Database Error")
	}

	return nil
}

func (c *commentPG) DeleteCommentById(commentId int) errs.MessageErr {
	_, err := c.db.Exec("DELETE FROM comments WHERE id = $1", commentId)

	if err != nil {
		return errs.NewInternalServerError("Database Error")
	}

	return nil 
}