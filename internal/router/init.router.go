package router

import (
	"github.com/labstack/echo"
	ihandler "github.com/nuttchai/go-rest/internal/handler/interface"
)

func Register(e *echo.Echo, sampleHandler ihandler.ISampleHandler, userHandler ihandler.IUserHandler) {
	initUserRouter(e, userHandler)
	initSampleRouter(e, sampleHandler)
}
