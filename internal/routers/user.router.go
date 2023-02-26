package routers

import (
	"github.com/labstack/echo"
	"github.com/nuttchai/go-rest/internal/handlers"
	"github.com/nuttchai/go-rest/internal/utils/api"
)

func InitUserRouter(e *echo.Echo) {
	handlers.InitUserHandler()

	e.GET(api.CreatePath("user/:id"), handlers.UserHandler.GetUser)
}
