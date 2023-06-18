package photo_pg

import (
	"database/sql"
	"final-project-2/entity"
	"final-project-2/pkg/errs"
	"final-project-2/repository/photo_repository"
	"fmt"
)

const (
	createNewPhoto = `
		INSERT INTO "photos"
		(
			user_id,
			title,
			caption,
			photo_url
		)
		VALUES ($1, $2, $3, $4);
	`

	getAllPhotoByUserId = `
		SELECT * from photos
		WHERE user_id = $1;
	`

	updatePhotoByIdQuery = `
		UPDATE photos
		SET
			title = $1,
			caption = $2,
			photo_url = $3,
			updated_at = current_timestamp
		WHERE id = $4
		RETURNING id, title, caption, photo_url, user_id, created_at, updated_at;
	
	`
)

type photoPG struct {
	db *sql.DB
}

func NewPhotoPG(db *sql.DB) photo_repository.PhotoRepository {
	return &photoPG{
		db: db,
	}
}

func (p *photoPG) CreateNewPhoto(photo entity.Photo) errs.MessageErr {
	_, err := p.db.Exec(createNewPhoto, photo.UserId, photo.Title, photo.Caption, photo.PhotoUrl)
	if err != nil {
		return errs.NewInternalServerError("Database Error 1")
	}

	return nil
}

func (p *photoPG) GetAllPhotoByUserId(userId int) ([]entity.Photo, errs.MessageErr) {
	// Execute query
	rows, err := p.db.Query(getAllPhotoByUserId, userId)
	if err != nil {
		return nil, errs.NewInternalServerError("Database Error 2")
	}
	defer rows.Close()

	// Create slice of photo
	var photos []entity.Photo

	// Loop through the rows
	for rows.Next() {
		var photo entity.Photo

		// Scan each row
		err := rows.Scan(&photo.Id, &photo.Title, &photo.Caption, &photo.PhotoUrl, &photo.UserId, &photo.CreatedAt, &photo.UpdatedAt)
		fmt.Println(err)
		if err != nil {
			return nil, errs.NewInternalServerError("Database Error 3")
		}

		// Append to slice
		photos = append(photos, photo)
	}

	// Check if there is an error when looping
	if err = rows.Err(); err != nil {
		return nil, errs.NewInternalServerError("Database Error 4")
	}

	return photos, nil
}

//updatePhoto
func (p *photoPG) UpdatePhotoById(payload entity.Photo) errs.MessageErr {
	_, err := p.db.Exec(updatePhotoByIdQuery, payload.Title, payload.Caption, payload.PhotoUrl, payload.Id)

	fmt.Println(updatePhotoByIdQuery, payload.Title, payload.Caption, payload.PhotoUrl, payload.Id)

	if err != nil {

		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}

//deletePhoto
func (p *photoPG) DeletePhotoById(photoId int) errs.MessageErr {
	_, err := p.db.Exec("DELETE FROM photos WHERE id = $1", photoId)

	if err != nil {
		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}