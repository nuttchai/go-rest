package handler

import (
	"github.com/labstack/echo"
	"github.com/nuttchai/go-rest/internal/constant"
	ihandler "github.com/nuttchai/go-rest/internal/handler/interface"
	iservice "github.com/nuttchai/go-rest/internal/service/interface"
	"github.com/nuttchai/go-rest/internal/utils/api"
)

type TUserHandler struct {
	userService iservice.IUserService
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

	res := api.SuccessResponse(user, constant.GetUserSuccessMsg)
	return c.JSON(res.Status, res)
}
