package interfaces

import (
	"github.com/nuttchai/go-rest/internal/models"
)

type IUserRepository interface {
	GetUser(id string) (*models.User, error)
}
