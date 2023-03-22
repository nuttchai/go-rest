package router

import "github.com/labstack/echo"

func Init(e *echo.Echo) {
	initUserRouter(e)
	initSampleRouter(e)
}
