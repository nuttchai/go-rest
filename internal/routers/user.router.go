package routers

import (
	"github.com/labstack/echo"
	"github.com/nuttchai/go-rest/internal/handlers"
	"github.com/nuttchai/go-rest/internal/utils/api"
)

func InitUserRouter(e *echo.Echo) {
	handler := handlers.InitUserHandler()

	e.GET(api.CreatePath("user/:id"), handler.GetUser)
}
