package service

import (
	"github.com/nuttchai/go-rest/internal/model"
	irepository "github.com/nuttchai/go-rest/internal/repository/interface"
	"github.com/nuttchai/go-rest/internal/service/interfaces"
)

type TUserService struct {
	repository irepository.IUserRepository
}

var (
	UserService interfaces.IUserService
)

func (s *TUserService) GetUser(id string) (*model.User, error) {
	return s.repository.RetrieveOne(id)
}
