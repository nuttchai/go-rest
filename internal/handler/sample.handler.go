package handler

import (
	"strconv"

	"github.com/labstack/echo"
	"github.com/nuttchai/go-rest/internal/constant"
	sampledto "github.com/nuttchai/go-rest/internal/dto/sample"
	shareddto "github.com/nuttchai/go-rest/internal/dto/shared"
	ihandler "github.com/nuttchai/go-rest/internal/handler/interface"
	iservice "github.com/nuttchai/go-rest/internal/service/interface"
	"github.com/nuttchai/go-rest/internal/util/api"
	jsonGen "github.com/nuttchai/go-rest/internal/util/json"
)

type TSampleHandler struct {
	sampleService iservice.ISampleService
	userService   iservice.IUserService
}

var (
	SampleHandler ihandler.ISampleHandler
)

func (h *TSampleHandler) Test(c echo.Context) error {
	resultTest := h.sampleService.Test()
	res := api.SuccessResponse(resultTest, constant.TestSuccessMsg)
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
	res := api.SuccessResponse(result, constant.GetSampleSuccessMsg)
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

	res := api.SuccessResponse(createdSample, constant.CreateSampleSuccessMsg)
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
		jsonErr := jsonGen.GenerateNotFoundIfErrorMatched(err, constant.SampleNotFound)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	res := api.SuccessResponse(updatedSample, constant.UpdateSampleSuccessMsg)
	return c.JSON(res.Status, res)
}

func (h *TSampleHandler) DeleteSample(c echo.Context) error {
	id := c.Param("id")
	err := h.sampleService.DeleteSample(id)
	if err != nil {
		jsonErr := jsonGen.GenerateNotFoundIfErrorMatched(err, constant.SampleNotFound)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	json := &shareddto.ValidatorResultDTO{Action: "DeleteSample", IsSuccess: true}
	res := api.SuccessResponse(json, constant.DeleteSampleSuccessMsg)
	return c.JSON(res.Status, res)
}
