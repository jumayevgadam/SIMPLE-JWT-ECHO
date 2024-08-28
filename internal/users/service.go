package users

import (
	userModel "SIMPLE-JWT-ECHO/internal/models/user"
	"context"
)

// Service is
type Service interface {
	SignUp(ctx context.Context, userReq *userModel.UserReq) (int, error)
	// Login(ctx context.Context, userReq *userModel.UserReq) (string, error)
}
