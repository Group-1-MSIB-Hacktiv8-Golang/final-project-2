package user_pg

import (
	"database/sql"
	"errors"
	"final-project-2/entity"
	"final-project-2/pkg/errs"
	"final-project-2/repository/user_repository"
	"fmt"
)

const (
	createNewUser = `
		INSERT INTO "users"
		(
			email,
			password,
			username,
			age
		)
		VALUES ($1, $2, $3, $4)
	`
	retrieveUserByEmail = `
		SELECT id, email, password from users
		WHERE email = $1;
	`

	
	retrieveUserById = `
		SELECT id, email, password from users
		WHERE id = $1;
	`
)

type userPG struct {
	db *sql.DB
}

func NewUserPG(db *sql.DB) user_repository.UserRepository {
	return &userPG{
		db: db,
	}
}

func (u *userPG) CreateNewUser(user entity.User) errs.MessageErr {
	_, err := u.db.Exec(createNewUser, user.Email, user.Password, user.Username, user.Age)

	if err != nil {
		return errs.NewInternalServerError(fmt.Errorf("failed to create new user: %w", err).Error())
	}

	return nil
}

func (u *userPG) GetUserById(userId int) (*entity.User, errs.MessageErr) {
	var user entity.User

	row := u.db.QueryRow(retrieveUserById, userId)

	err := row.Scan(&user.Id, &user.Email, &user.Password)

	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, errs.NewNotFoundError("user not found")
		}
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &user, nil
}

func (u *userPG) GetUserByEmail(userEmail string) (*entity.User, errs.MessageErr) {
	var user entity.User

	row := u.db.QueryRow(retrieveUserByEmail, userEmail)

	err := row.Scan(&user.Id, &user.Email, &user.Password)

	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, errs.NewNotFoundError("user not found")
		}
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &user, nil
}
