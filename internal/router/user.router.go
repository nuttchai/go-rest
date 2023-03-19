package router

import (
	"github.com/labstack/echo"
	handler "github.com/nuttchai/go-rest/internal/handler"
	"github.com/nuttchai/go-rest/internal/util/api"
)

func initUserRouter(e *echo.Echo) {
	e.GET(api.CreatePath("user/:id"), handler.UserHandler.GetUser)
}
