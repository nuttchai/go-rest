package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/nuttchai/go-rest/internal/constants"
	sampledto "github.com/nuttchai/go-rest/internal/dto/sample"
	"github.com/nuttchai/go-rest/internal/models"
	"github.com/stretchr/testify/assert"
)

var (
	e                *echo.Echo
	testMock         func() string
	getSampleMock    func(id string) (*models.Sample, error)
	createSampleMock func(sample *sampledto.CreateSampleDTO) (*models.Sample, error)
	updateSampleMock func(sample *sampledto.UpdateSampleDTO) (*models.Sample, error)
	deleteSampleMock func(id string) error
	getUserMock      func(id string) (*models.User, error)
)

type TSampleServiceMock struct{}
type TUserServiceMock struct{}

func (*TSampleServiceMock) Test() string {
	return testMock()
}

func (*TSampleServiceMock) GetSample(id string) (*models.Sample, error) {
	return getSampleMock(id)
}

func (*TSampleServiceMock) CreateSample(sample *sampledto.CreateSampleDTO) (*models.Sample, error) {
	return createSampleMock(sample)
}

func (*TSampleServiceMock) UpdateSample(sample *sampledto.UpdateSampleDTO) (*models.Sample, error) {
	return updateSampleMock(sample)
}

func (*TSampleServiceMock) DeleteSample(id string) error {
	return deleteSampleMock(id)
}

func (*TUserServiceMock) GetUser(id string) (*models.User, error) {
	return getUserMock(id)
}

func init() {
	e = echo.New()
}

func setUpRequest(method string, subPath string) *http.Request {
	path := constants.LocalHost + constants.BasePath + subPath
	req := httptest.NewRequest(method, path, nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	return req
}

func initSampleServiceMock() {
	SampleHandler = &TSampleHandler{
		sampleService: &TSampleServiceMock{},
		userService:   &TUserServiceMock{},
	}
}

func TestTestReturn(t *testing.T) {
	// Arrange
	testMock = func() string {
		return "test"
	}
	initSampleServiceMock()

	rec := httptest.NewRecorder()
	req := setUpRequest(echo.GET, "/sample")
	c := e.NewContext(req, rec)

	// Act
	err := SampleHandler.Test(c)

	// Assert
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusOK, rec.Code)
}

func TestGetSampleReturn(t *testing.T) {
	// Arrange
	getSampleMock = func(id string) (*models.Sample, error) {
		return &models.Sample{
			Id:          1,
			Name:        "sample",
			Description: "description",
			OwnerId:     1,
		}, nil
	}
	getUserMock = func(id string) (*models.User, error) {
		return &models.User{
			Id:       1,
			Username: "username",
		}, nil
	}
	initSampleServiceMock()

	rec := httptest.NewRecorder()
	req := setUpRequest(echo.GET, "/sample/1")

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
	getSampleMock = func(id string) (*models.Sample, error) {
		return nil, errors.New("error")
	}
	initSampleServiceMock()

	rec := httptest.NewRecorder()
	req := setUpRequest(echo.GET, "/sample/1")

	c := e.NewContext(req, rec)

	// Act
	err := SampleHandler.GetSample(c)

	// Assert
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, rec.Code)
}

func TestGetSampleReturnErrorFromGetUser(t *testing.T) {
	// Arrange
	getSampleMock = func(id string) (*models.Sample, error) {
		return &models.Sample{
			Id:          1,
			Name:        "sample",
			Description: "description",
			OwnerId:     1,
		}, nil
	}
	getUserMock = func(id string) (*models.User, error) {
		return nil, errors.New("error")
	}
	initSampleServiceMock()

	rec := httptest.NewRecorder()
	req := setUpRequest(echo.GET, "/sample/1")

	c := e.NewContext(req, rec)

	// Act
	err := SampleHandler.GetSample(c)

	// Assert
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, rec.Code)
}