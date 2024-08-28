package repository

import (
	"SIMPLE-JWT-ECHO/internal/connection"
	userModel "SIMPLE-JWT-ECHO/internal/models/user"
	"context"
)

// UserRepository is
type UserRepository struct {
	psqlDB connection.DB
}

// NewUserRepository is
func NewUserRepository(psqlDB connection.DB) *UserRepository {
	return &UserRepository{psqlDB: psqlDB}
}

// SignUp is
func (ur *UserRepository) SignUp(ctx context.Context, userDAO *userModel.UserRes) (int, error) {
	var userID int

	err := ur.psqlDB.QueryRow(
		ctx,
		signUPQuery,
		userDAO.Username,
		userDAO.Email,
		userDAO.Password,
	).Scan(&userID)

	if err != nil {
		return -1, ErrSignUpRepo
	}

	return userID, nil
}
