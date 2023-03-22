package mhttp

import (
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo"
	"github.com/nuttchai/go-rest/internal/constant"
)

func SetupMockRequest(method string, subPath string) *http.Request {
	path := constant.LocalHost + constant.BasePath + subPath
	req := httptest.NewRequest(method, path, nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	return req
}
