package routers

import (
	"github.com/labstack/echo"
	"github.com/nuttchai/go-rest/internal/handlers"
	"github.com/nuttchai/go-rest/internal/utils/api"
)

func InitSampleRouter(e *echo.Echo) *echo.Echo {
	handler := handlers.InitSampleHandler()

	e.GET(api.CreatePath("sample"), handler.Test)
	e.GET(api.CreatePath("sample/:id"), handler.GetSample)

	e.POST(api.CreatePath("sample"), handler.CreateSample)

	e.PUT(api.CreatePath("sample/:id"), handler.UpdateSample)

	e.DELETE(api.CreatePath("sample/:id"), handler.DeleteSample)
	return e
}
