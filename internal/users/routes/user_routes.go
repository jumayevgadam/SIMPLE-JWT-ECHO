package routes

import (
	"SIMPLE-JWT-ECHO/internal/users"

	"github.com/labstack/echo/v4"
)

// MapUserRoutes is
func MapUserRoutes(userGroup *echo.Group, userHandler users.Handler) {
	userGroup.POST("/sign-up", userHandler.SignUp())
}
