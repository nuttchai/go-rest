package interfaces

import (
	"github.com/nuttchai/go-rest/internal/models"
)

type IUserRepository interface {
	RetrieveOne(id string) (*models.User, error)
}
