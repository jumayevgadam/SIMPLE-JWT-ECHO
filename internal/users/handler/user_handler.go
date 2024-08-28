package handler

import (
	userModel "SIMPLE-JWT-ECHO/internal/models/user"
	userServInterface "SIMPLE-JWT-ECHO/internal/users"

	"github.com/labstack/echo/v4"
)

// UserHandler is
type UserHandler struct {
	service userServInterface.Service
}

// NewUserHandler is
func NewUserHandler(service userServInterface.Service) *UserHandler {
	return &UserHandler{service: service}
}

// SignUp is
func (uh *UserHandler) SignUp() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get the req body
		var userReq userModel.UserReq
		if c.Bind(&userReq) != nil {
			return c.JSON(
				400,
				"error bind user",
			)
		}

		// Create The User, call service
		userResp, err := uh.service.SignUp(c.Request().Context(), &userReq)
		if err != nil {
			return c.JSON(
				500,
				err.Error())
		}

		return c.JSON(
			200,
			userResp,
		)
	}
}
