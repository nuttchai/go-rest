package service

import (
	"github.com/nuttchai/go-rest/internal/model"
	irepository "github.com/nuttchai/go-rest/internal/repository/interface"
	iservice "github.com/nuttchai/go-rest/internal/service/interface"
)

type TUserService struct {
	Repository irepository.IUserRepository
}

var (
	UserService iservice.IUserService
)

func initUserService(userService *TUserService) {
	UserService = userService
}

func (s *TUserService) GetUser(id string) (*model.User, error) {
	return s.Repository.RetrieveOne(id)
}
