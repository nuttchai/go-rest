package handler

import (
	"github.com/labstack/echo"
	"github.com/nuttchai/go-rest/internal/constants"
	"github.com/nuttchai/go-rest/internal/handler/interfaces"
	servicesInterfaces "github.com/nuttchai/go-rest/internal/services/interfaces"
	"github.com/nuttchai/go-rest/internal/utils/api"
)

type TUserHandler struct {
	userService servicesInterfaces.IUserService
}

var (
	UserHandler interfaces.IUserHandler
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
