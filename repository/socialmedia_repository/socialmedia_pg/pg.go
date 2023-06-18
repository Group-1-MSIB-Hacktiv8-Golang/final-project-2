package socialmedia_pg

import (
	"database/sql"
	"final-project-2/entity"
	"final-project-2/pkg/errs"
	"final-project-2/repository/socialmedia_repository"
	"fmt"
)

const (
	createNewSocialMedia = `
		INSERT INTO "social_media"
		(
			user_id,
			name,
			social_media_url
		)
		VALUES ($1, $2, $3);
	`

	getAllSocialMediaByUserId = `
		SELECT * from social_media
		WHERE user_id = $1;
	`

	updateSocialMediaByIdQuery = `
		UPDATE social_media
		SET
			name = $1,
			social_media_url = $2,
			updated_at = current_timestamp
		WHERE id = $3
		RETURNING id, name, social_media_url, user_id, created_at, updated_at;
	`
)

type socialMediaPG struct {
	db *sql.DB
}

func NewSocialMediaPG(db *sql.DB) socialmedia_repository.SocialMediaRepository {
	return &socialMediaPG{
		db: db,
	}
}

func (s *socialMediaPG) CreateNewSocialMedia(socialMedia entity.SocialMedia) errs.MessageErr {
	_, err := s.db.Exec(createNewSocialMedia, socialMedia.UserId, socialMedia.Name, socialMedia.SocialMediaUrl)
	if err != nil {
		return errs.NewInternalServerError("Database Error")
	}

	return nil
}

func (s *socialMediaPG) GetAllSocialMediaByUserId(userId int) ([]entity.SocialMedia, errs.MessageErr) {
	// Execute query
	rows, err := s.db.Query(getAllSocialMediaByUserId, userId)
	if err != nil {
		return nil, errs.NewInternalServerError("Database Error 2")
	}
	defer rows.Close()

	// Create slice of socialMedia
	var socialMedias []entity.SocialMedia

	// Loop through the rows
	for rows.Next() {
		var socialMedia entity.SocialMedia

		// Scan each row
		err := rows.Scan(&socialMedia.Id, &socialMedia.Name, &socialMedia.SocialMediaUrl, &socialMedia.UserId)
		fmt.Println(err)
		if err != nil {
			return nil, errs.NewInternalServerError("Database Error 3")
		}

		// Append to slice
		socialMedias = append(socialMedias, socialMedia)
	}

	// Check if there is an error when looping
	if err = rows.Err(); err != nil {
		return nil, errs.NewInternalServerError("Database Error 4")
	}

	return socialMedias, nil
}

func (p *socialMediaPG) UpdateSocialMediaById(payload entity.SocialMedia) errs.MessageErr {
	_, err := p.db.Exec(updateSocialMediaByIdQuery, payload.Name, payload.SocialMediaUrl, payload.Id)

	fmt.Println(err)

	if err != nil {

		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}

func (p *socialMediaPG) DeleteSocialMediaById(socialMediaId int) errs.MessageErr {
	_, err := p.db.Exec("DELETE FROM social_media WHERE id = $1", socialMediaId)

	if err != nil {
		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}