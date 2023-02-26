package routers

import (
	"github.com/labstack/echo"
	"github.com/nuttchai/go-rest/internal/handlers"
	"github.com/nuttchai/go-rest/internal/utils/api"
)

func initUserRouter(e *echo.Echo) {
	e.GET(api.CreatePath("user/:id"), handlers.UserHandler.GetUser)
}
