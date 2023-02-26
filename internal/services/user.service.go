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

func InitUserService() IUserService {
	UserService = &TUserService{
		repository: repositories.InitUserRepository(),
	}
	return UserService
}

func (s *TUserService) GetUser(id string) (*models.User, error) {
	return s.repository.GetUser(id)
}
