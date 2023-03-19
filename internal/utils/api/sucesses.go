package api

import (
	"net/http"

	"github.com/nuttchai/go-rest/internal/constant"
)

type APISuccess struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

func SuccessResponse(result interface{}, message ...string) *APISuccess {
	msg := constant.DefaultAPISuccessMsg
	if len(message) > 0 {
		msg = message[0]
	}

	return &APISuccess{
		Status:  http.StatusOK,
		Message: msg,
		Result:  result,
	}
}
