package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	mhandler "github.com/nuttchai/go-rest/internal/handler/mock"
	"github.com/nuttchai/go-rest/internal/model"
	mhttp "github.com/nuttchai/go-rest/internal/util/mock"
	"github.com/stretchr/testify/assert"
)

var (
	e           *echo.Echo
	handlerMock *TSampleHandler
)

func init() {
	e = echo.New()
	handlerMock = &TSampleHandler{
		SampleService: &mhandler.TSampleServiceMock{},
		UserService:   &mhandler.TUserServiceMock{},
	}
}

func TestTestReturn(t *testing.T) {
	// Arrange
	mhandler.TestMock = func() string {
		return "test"
	}

	initSampleHandler(handlerMock)

	rec := httptest.NewRecorder()
	req := mhttp.SetupMockRequest(echo.GET, "/sample")
	c := e.NewContext(req, rec)

	// Act
	err := SampleHandler.Test(c)

	// Assert
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusOK, rec.Code)
}

func TestGetSampleReturn(t *testing.T) {
	// Arrange
	mhandler.GetSampleMock = func(id string) (*model.Sample, error) {
		return &model.Sample{
			Id:          1,
			Name:        "sample",
			Description: "description",
			OwnerId:     1,
		}, nil
	}

	mhandler.GetUserMock = func(id string) (*model.User, error) {
		return &model.User{
			Id:       1,
			Username: "username",
		}, nil
	}

	initSampleHandler(handlerMock)

	rec := httptest.NewRecorder()
	req := mhttp.SetupMockRequest(echo.GET, "/sample/1")

	c := e.NewContext(req, rec)

	// Act
	err := SampleHandler.GetSample(c)

	// Assert
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusOK, rec.Code)
	assert.EqualValues(t, "{\"status\":200,\"message\":\"Get sample successfully\",\"result\":{\"id\":1,\"name\":\"sample\",\"description\":\"description\",\"owner\":\"username\"}}\n",
		rec.Body.String())
}

func TestGetSampleReturnErrorFromGetSample(t *testing.T) {
	// Arrange
	mhandler.GetSampleMock = func(id string) (*model.Sample, error) {
		return nil, errors.New("error")
	}

	initSampleHandler(handlerMock)

	rec := httptest.NewRecorder()
	req := mhttp.SetupMockRequest(echo.GET, "/sample/1")

	c := e.NewContext(req, rec)

	// Act
	err := SampleHandler.GetSample(c)

	// Assert
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, rec.Code)
}

func TestGetSampleReturnErrorFromGetUser(t *testing.T) {
	// Arrange
	mhandler.GetSampleMock = func(id string) (*model.Sample, error) {
		return &model.Sample{
			Id:          1,
			Name:        "sample",
			Description: "description",
			OwnerId:     1,
		}, nil
	}

	mhandler.GetUserMock = func(id string) (*model.User, error) {
		return nil, errors.New("error")
	}

	initSampleHandler(handlerMock)

	rec := httptest.NewRecorder()
	req := mhttp.SetupMockRequest(echo.GET, "/sample/1")
	c := e.NewContext(req, rec)

	// Act
	err := SampleHandler.GetSample(c)

	// Assert
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, rec.Code)
}
