package handler

import (
	"github.com/labstack/echo"
	"github.com/nuttchai/go-rest/internal/constants"
	ihandler "github.com/nuttchai/go-rest/internal/handler/interface"
	ihandlersvc "github.com/nuttchai/go-rest/internal/services/interfaces"
	"github.com/nuttchai/go-rest/internal/utils/api"
)

type TUserHandler struct {
	userService ihandlersvc.IUserService
}

var (
	UserHandler ihandler.IUserHandler
)

func (h *TUserHandler) GetUser(c echo.Context) error {
	id := c.Param("id")
	user, err := h.userService.GetUser(id)
	if err != nil {
		jsonErr := api.InternalServerError(err)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	res := api.SuccessResponse(user, constants.GetUserSuccessMsg)
	return c.JSON(res.Status, res)
}
