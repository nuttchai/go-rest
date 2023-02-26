package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	sampledto "github.com/nuttchai/go-rest/internal/dto/sample"
	"github.com/nuttchai/go-rest/internal/models"
	"github.com/stretchr/testify/assert"
)

var (
	e                *echo.Echo
	req              *http.Request
	testMock         func() string
	getSampleMock    func(id string) (*models.Sample, error)
	createSampleMock func(sample *sampledto.CreateSampleDTO) (*models.Sample, error)
	updateSampleMock func(sample *sampledto.UpdateSampleDTO) (*models.Sample, error)
	deleteSampleMock func(id string) error

	getUserMock func(id string) (*models.User, error)
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
	req = httptest.NewRequest(echo.GET, "http://localhost:8000", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
}

func InitSampleServiceMock() {
	SampleHandler = &TSampleHandler{
		sampleService: &TSampleServiceMock{},
		userService:   &TUserServiceMock{},
	}
}

func TestSampleHandlerTestReturn(t *testing.T) {
	// Arrange
	testMock = func() string {
		return "test"
	}
	InitSampleServiceMock()

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/sample")

	// Act
	err := SampleHandler.Test(c)

	// Assert
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusOK, rec.Code)
}

func TestSampleHandlerGetSampleReturn(t *testing.T) {
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

	InitSampleServiceMock()

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/sample/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Act
	err := SampleHandler.GetSample(c)

	// Assert
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusOK, rec.Code)
	assert.EqualValues(t, "{\"status\":200,\"message\":\"Get sample successfully\",\"result\":{\"id\":1,\"name\":\"sample\",\"description\":\"description\",\"owner\":\"username\"}}\n",
		rec.Body.String())
}
