package services

import (
	"github.com/nuttchai/go-rest/internal/models"
	"github.com/nuttchai/go-rest/internal/repositories"
)

type TUserService struct {
	repository repositories.IUserRepository
}

var (
	UserService IUserService
)

func (s *TUserService) GetUser(id string) (*models.User, error) {
	return s.repository.GetUser(id)
}
