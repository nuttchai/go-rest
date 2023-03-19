package router

import (
	"github.com/labstack/echo"
	handler "github.com/nuttchai/go-rest/internal/handler"
	"github.com/nuttchai/go-rest/internal/util/api"
)

func initSampleRouter(e *echo.Echo) {
	e.GET(api.CreatePath("sample"), handler.SampleHandler.Test)
	e.GET(api.CreatePath("sample/:id"), handler.SampleHandler.GetSample)

	e.POST(api.CreatePath("sample"), handler.SampleHandler.CreateSample)

	e.PUT(api.CreatePath("sample"), handler.SampleHandler.UpdateSample)

	e.DELETE(api.CreatePath("sample/:id"), handler.SampleHandler.DeleteSample)
}
