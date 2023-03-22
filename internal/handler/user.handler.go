package handler

import (
	"github.com/labstack/echo"
	"github.com/nuttchai/go-rest/internal/constant"
	ihandler "github.com/nuttchai/go-rest/internal/handler/interface"
	iservice "github.com/nuttchai/go-rest/internal/service/interface"
	"github.com/nuttchai/go-rest/internal/util/api"
)

type TUserHandler struct {
	UserService iservice.IUserService
}

var (
	UserHandler ihandler.IUserHandler
)

func initUserHandler(userHandler *TUserHandler) {
	UserHandler = userHandler
}

func (h *TUserHandler) GetUser(c echo.Context) error {
	id := c.Param("id")
	user, err := h.UserService.GetUser(id)
	if err != nil {
		jsonErr := api.InternalServerError(err)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	res := api.SuccessResponse(user, constant.GetUserSuccessMsg)
	return c.JSON(res.Status, res)
}
