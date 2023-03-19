package services

import (
	"github.com/nuttchai/go-rest/internal/model"
	repositoriesInterfaces "github.com/nuttchai/go-rest/internal/repository/interfaces"
	"github.com/nuttchai/go-rest/internal/services/interfaces"
)

type TUserService struct {
	repository repositoriesInterfaces.IUserRepository
}

var (
	UserService interfaces.IUserService
)

func (s *TUserService) GetUser(id string) (*model.User, error) {
	return s.repository.RetrieveOne(id)
}
