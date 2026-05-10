package router

import (
	"github.com/labstack/echo"
	ihandler "github.com/nuttchai/go-rest/internal/handler/interface"
	"github.com/nuttchai/go-rest/internal/util/api"
)

func initUserRouter(e *echo.Echo, userHandler ihandler.IUserHandler) {
	e.GET(api.CreatePath("user/:id"), userHandler.GetUser)
}
