package services

import (
	"github.com/nuttchai/go-rest/internal/models"
	repositoriesInterfaces "github.com/nuttchai/go-rest/internal/repositories/interfaces"
	"github.com/nuttchai/go-rest/internal/services/interfaces"
)

type TUserService struct {
	repository repositoriesInterfaces.IUserRepository
}

var (
	UserService interfaces.IUserService
)

func (s *TUserService) GetUser(id string) (*models.User, error) {
	return s.repository.GetUser(id)
}
