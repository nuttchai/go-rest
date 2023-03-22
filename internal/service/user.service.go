package service

import (
	"github.com/nuttchai/go-rest/internal/model"
	"github.com/nuttchai/go-rest/internal/repository"
	irepository "github.com/nuttchai/go-rest/internal/repository/interface"
	iservice "github.com/nuttchai/go-rest/internal/service/interface"
)

type TUserService struct {
	Repository irepository.IUserRepository
}

var (
	UserService iservice.IUserService
)

func initUserService() {
	UserService = &TUserService{
		Repository: repository.UserRepository,
	}
}

func (s *TUserService) GetUser(id string) (*model.User, error) {
	return s.Repository.RetrieveOne(id)
}
