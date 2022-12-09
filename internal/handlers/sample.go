package handlers

import (
	"github.com/labstack/echo"
	"github.com/nuttchai/go-rest/internal/constants"
	sampledto "github.com/nuttchai/go-rest/internal/dto/sample"
	shareddto "github.com/nuttchai/go-rest/internal/dto/shared"
	"github.com/nuttchai/go-rest/internal/services"
	"github.com/nuttchai/go-rest/internal/utils/api"
	jsonGen "github.com/nuttchai/go-rest/internal/utils/json"
)

type sampleHandler struct{}

type sampleHandlerInterface interface {
	GetSample(c echo.Context) error
	CreateSample(c echo.Context) error
	UpdateSample(c echo.Context) error
	DeleteSample(c echo.Context) error
}

var (
	SampleHandler sampleHandlerInterface
)

func init() {
	SampleHandler = &sampleHandler{}
}

func (h *sampleHandler) GetSample(c echo.Context) error {
	id := c.Param("id")
	sample, err := services.SampleService.GetSample(id)
	if err != nil {
		jsonErr := api.InternalServerError(err)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	res := api.SuccessResponse(sample, constants.GetSampleSuccessMsg)
	return c.JSON(res.Status, res)
}

func (h *sampleHandler) CreateSample(c echo.Context) error {
	var sampleDto *sampledto.CreateSampleDTO
	if err := api.DecodeDTO(c, &sampleDto); err != nil {
		jsonErr := api.BadRequestError(err)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	createdSample, err := services.SampleService.CreateSample(sampleDto)
	if err != nil {
		jsonErr := api.InternalServerError(err)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	res := api.SuccessResponse(createdSample, constants.CreateSampleSuccessMsg)
	return c.JSON(res.Status, res)
}

func (h *sampleHandler) UpdateSample(c echo.Context) error {
	var sampleDto sampledto.UpdateSampleDTO
	if err := api.DecodeDTO(c, &sampleDto); err != nil {
		jsonErr := api.BadRequestError(err)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	updatedSample, err := services.SampleService.UpdateSample(&sampleDto)
	if err != nil {
		jsonErr := jsonGen.GenerateNotFoundIfErrorMatched(err, constants.SampleNotFound)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	res := api.SuccessResponse(updatedSample, constants.UpdateSampleSuccessMsg)
	return c.JSON(res.Status, res)
}

func (h *sampleHandler) DeleteSample(c echo.Context) error {
	id := c.Param("id")
	err := services.SampleService.DeleteSample(id)
	if err != nil {
		jsonErr := jsonGen.GenerateNotFoundIfErrorMatched(err, constants.SampleNotFound)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	json := &shareddto.ValidatorResultDTO{Action: "DeleteSample", IsSuccess: true}
	res := api.SuccessResponse(json, constants.DeleteSampleSuccessMsg)
	return c.JSON(res.Status, res)
}
