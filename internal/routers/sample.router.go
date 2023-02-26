package routers

import (
	"github.com/labstack/echo"
	"github.com/nuttchai/go-rest/internal/handlers"
	"github.com/nuttchai/go-rest/internal/utils/api"
)

func initSampleRouter(e *echo.Echo) {
	e.GET(api.CreatePath("sample"), handlers.SampleHandler.Test)
	e.GET(api.CreatePath("sample/:id"), handlers.SampleHandler.GetSample)

	e.POST(api.CreatePath("sample"), handlers.SampleHandler.CreateSample)

	e.PUT(api.CreatePath("sample"), handlers.SampleHandler.UpdateSample)

	e.DELETE(api.CreatePath("sample/:id"), handlers.SampleHandler.DeleteSample)
}
