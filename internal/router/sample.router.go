package router

import (
	"github.com/labstack/echo"
	ihandler "github.com/nuttchai/go-rest/internal/handler/interface"
	"github.com/nuttchai/go-rest/internal/util/api"
)

func initSampleRouter(e *echo.Echo, sampleHandler ihandler.ISampleHandler) {
	e.GET(api.CreatePath("sample"), sampleHandler.Test)
	e.GET(api.CreatePath("sample/:id"), sampleHandler.GetSample)

	e.POST(api.CreatePath("sample"), sampleHandler.CreateSample)

	e.PUT(api.CreatePath("sample"), sampleHandler.UpdateSample)

	e.DELETE(api.CreatePath("sample/:id"), sampleHandler.DeleteSample)
}
