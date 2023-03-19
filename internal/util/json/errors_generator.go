package json

import (
	"github.com/nuttchai/go-rest/internal/util/api"
)

func GenerateNotFoundIfErrorMatched(err error, expectedErr string) *api.APIError {
	if err.Error() == expectedErr {
		return api.NotFoundError(err)
	}
	return api.InternalServerError(err)
}
