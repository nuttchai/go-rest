package handler

import (
	"github.com/labstack/echo"
	"github.com/nuttchai/go-rest/internal/constant"
	iservice "github.com/nuttchai/go-rest/internal/service/interface"
	"github.com/nuttchai/go-rest/internal/util/api"
)

type TUserHandler struct {
	UserService iservice.IUserService
}

func NewUserHandler(userService iservice.IUserService) *TUserHandler {
	return &TUserHandler{
		UserService: userService,
	}
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
