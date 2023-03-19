package api

import (
	"net/http"

	"github.com/nuttchai/go-rest/internal/constant"
)

type APIError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

func CustomError(err error, status int, message ...string) *APIError {
	msg := constant.DefaultAPIErrorMsg
	if len(message) > 0 {
		msg = message[0]
	}

	return &APIError{
		Status:  status,
		Message: msg,
		Error:   err.Error(),
	}
}

func InternalServerError(err error) *APIError {
	return &APIError{
		Status:  http.StatusInternalServerError,
		Message: "InternalServerError",
		Error:   err.Error(),
	}
}

func NotFoundError(err error) *APIError {
	return &APIError{
		Status:  http.StatusNotFound,
		Message: "NotFoundError",
		Error:   err.Error(),
	}
}

func BadRequestError(err error) *APIError {
	return &APIError{
		Status:  http.StatusBadRequest,
		Message: "BadRequestError",
		Error:   err.Error(),
	}
}
