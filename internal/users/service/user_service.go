package service

import (
	"SIMPLE-JWT-ECHO/internal/database"
	userModel "SIMPLE-JWT-ECHO/internal/models/user"
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// UserService is
type UserService struct {
	repo database.DataStore
}

// NewUserService is
func NewUserService(repo database.DataStore) *UserService {
	return &UserService{repo: repo}
}

// SignUp is
func (us *UserService) SignUp(ctx context.Context, userReq *userModel.UserReq) (int, error) {
	var userID int

	if err := us.repo.WithTransaction(ctx, func(db database.DataStore) error {
		// Hash the Password
		hashPass, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), 30)
		if err != nil {
			return fmt.Errorf("failed to hash password: %w", err)
		}
		userReq.Password = string(hashPass)

		userID, err = db.UsersRepo().SignUp(ctx, userReq.ToStorage())
		if err != nil {
			return fmt.Errorf("db.UsersRepo.SignUp: %v", err.Error())
		}

		return nil
	}); err != nil {
		return -1, fmt.Errorf("[service][SignUP][W-TX]: %w", err)
	}

	return userID, nil
}
