package service

import (
	"github.com/nuttchai/go-rest/internal/model"
	irepository "github.com/nuttchai/go-rest/internal/repository/interface"
	iservice "github.com/nuttchai/go-rest/internal/service/interface"
)

type TUserService struct {
	repository irepository.IUserRepository
}

var (
	UserService iservice.IUserService
)

func (s *TUserService) GetUser(id string) (*model.User, error) {
	return s.repository.RetrieveOne(id)
}
