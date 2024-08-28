package server

import (
	userHandler "SIMPLE-JWT-ECHO/internal/users/handler"
	userRoutes "SIMPLE-JWT-ECHO/internal/users/routes"
	userService "SIMPLE-JWT-ECHO/internal/users/service"

	"github.com/labstack/echo/v4"
)

const (
	userGroup = "/api/user"
)

// MapHandlers is
func (s *Server) MapHandlers(e *echo.Echo) {
	// Init Services
	UserService := userService.NewUserService(s.dataStore)

	// Init Handlers
	UserHandler := userHandler.NewUserHandler(UserService)

	// Init Main Group Routes
	UserGroup := s.echo.Group(userGroup)

	// Init Routes
	userRoutes.MapUserRoutes(UserGroup, UserHandler)
}
