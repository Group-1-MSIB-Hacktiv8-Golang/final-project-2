package user_pg

import (
	"database/sql"
	"errors"
	"final-project-2/entity"
	"final-project-2/pkg/errs"
	"final-project-2/repository/user_repository"
	"fmt"
	"time"
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
		VALUES ($1, $2, $3, $4);
	`
	retrieveUserByEmail = `
		SELECT id, email, password from users
		WHERE email = $1;
	`

	
	retrieveUserById = `
		SELECT * from users
		WHERE id = $1;
	`

	//update user with return and update_at now()
	updateUser = `
		UPDATE users
		SET email = $1, username = $2, updated_at = current_timestamp
		WHERE id = $3
		RETURNING id, email, username, age, created_at, updated_at;
	`

	//delete user
	deleteUser = `
		DELETE FROM users
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

	err := row.Scan(&user.Id, &user.Email, &user.Password, &user.Username, &user.Age, &user.CreatedAt, &user.UpdatedAt)

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


//UpdateUser
func (u *userPG) UpdateUser(payload *entity.User, userId int) (*entity.User, errs.MessageErr) {

	var user entity.User

	payload.UpdatedAt = time.Now()

	tx, err := u.db.Begin()

	if err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	fmt.Println("QueryRow di repo", updateUser, payload.Email, payload.Username, userId)

	row := u.db.QueryRow(updateUser, payload.Email, payload.Username, userId)

	err = row.Scan(&user.Id, &user.Email, &user.Username, &user.Age, &user.CreatedAt, &user.UpdatedAt)

	fmt.Println("err di repo", err)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, errs.NewNotFoundError("user not found")
		}
		return nil, errs.NewInternalServerError("failed to update user")
	}

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &user, nil
}

//DeleteUser
func (u *userPG) DeleteUser(userId int) errs.MessageErr {
	
	tx, err := u.db.Begin()

	if err != nil {
		return errs.NewInternalServerError("something went wrong")
	}

	_, err = u.db.Exec(deleteUser, userId)

	if err != nil {
		tx.Rollback()
		return errs.NewInternalServerError("something went wrong")
	}

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}