package services

import "github.com/nuttchai/go-rest/internal/models"

type IUserService interface {
	GetUser(id string) (*models.User, error)
}
