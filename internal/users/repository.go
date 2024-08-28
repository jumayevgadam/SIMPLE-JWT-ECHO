package users

import (
	userModel "SIMPLE-JWT-ECHO/internal/models/user"
	"context"
)

// Repository is
type Repository interface {
	SignUp(ctx context.Context, userDAO *userModel.UserRes) (int, error)
	// Login(ctx context.Context, userDAO *userModel.UserRes) (string, error)
}
