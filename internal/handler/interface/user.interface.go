package ihandler

import "github.com/labstack/echo"

type IUserHandler interface {
	GetUser(c echo.Context) error
}
