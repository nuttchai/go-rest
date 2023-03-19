package handler

import (
	"strconv"

	"github.com/labstack/echo"
	"github.com/nuttchai/go-rest/internal/constants"
	sampledto "github.com/nuttchai/go-rest/internal/dto/sample"
	shareddto "github.com/nuttchai/go-rest/internal/dto/shared"
	ihandler "github.com/nuttchai/go-rest/internal/handler/interface"
	ihandlersvc "github.com/nuttchai/go-rest/internal/service/interfaces"
	"github.com/nuttchai/go-rest/internal/utils/api"
	jsonGen "github.com/nuttchai/go-rest/internal/utils/json"
)

type TSampleHandler struct {
	sampleService ihandlersvc.ISampleService
	userService   ihandlersvc.IUserService
}

var (
	SampleHandler ihandler.ISampleHandler
)

func (h *TSampleHandler) Test(c echo.Context) error {
	resultTest := h.sampleService.Test()
	res := api.SuccessResponse(resultTest, constants.TestSuccessMsg)
	return c.JSON(res.Status, res)
}

func (h *TSampleHandler) GetSample(c echo.Context) error {
	id := c.Param("id")
	sample, err := h.sampleService.GetSample(id)
	if err != nil {
		jsonErr := api.InternalServerError(err)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	ownerId := strconv.Itoa(sample.OwnerId)
	owner, err := h.userService.GetUser(ownerId)
	if err != nil {
		jsonErr := api.InternalServerError(err)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	result := sampledto.GetSampleWithUserDTO{
		Id:          sample.Id,
		Name:        sample.Name,
		Description: sample.Description,
		Owner:       owner.Username,
	}
	res := api.SuccessResponse(result, constants.GetSampleSuccessMsg)
	return c.JSON(res.Status, res)
}

func (h *TSampleHandler) CreateSample(c echo.Context) error {
	var sampleDto *sampledto.CreateSampleDTO
	if err := api.DecodeDTO(c, &sampleDto); err != nil {
		jsonErr := api.BadRequestError(err)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	createdSample, err := h.sampleService.CreateSample(sampleDto)
	if err != nil {
		jsonErr := api.InternalServerError(err)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	res := api.SuccessResponse(createdSample, constants.CreateSampleSuccessMsg)
	return c.JSON(res.Status, res)
}

func (h *TSampleHandler) UpdateSample(c echo.Context) error {
	var sampleDto sampledto.UpdateSampleDTO
	if err := api.DecodeDTO(c, &sampleDto); err != nil {
		jsonErr := api.BadRequestError(err)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	updatedSample, err := h.sampleService.UpdateSample(&sampleDto)
	if err != nil {
		jsonErr := jsonGen.GenerateNotFoundIfErrorMatched(err, constants.SampleNotFound)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	res := api.SuccessResponse(updatedSample, constants.UpdateSampleSuccessMsg)
	return c.JSON(res.Status, res)
}

func (h *TSampleHandler) DeleteSample(c echo.Context) error {
	id := c.Param("id")
	err := h.sampleService.DeleteSample(id)
	if err != nil {
		jsonErr := jsonGen.GenerateNotFoundIfErrorMatched(err, constants.SampleNotFound)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	json := &shareddto.ValidatorResultDTO{Action: "DeleteSample", IsSuccess: true}
	res := api.SuccessResponse(json, constants.DeleteSampleSuccessMsg)
	return c.JSON(res.Status, res)
}
