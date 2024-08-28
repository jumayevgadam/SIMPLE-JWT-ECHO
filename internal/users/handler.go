package users

import "github.com/labstack/echo/v4"

// Handler is
type Handler interface {
	SignUp() echo.HandlerFunc
	//Login() echo.HandlerFunc
}
