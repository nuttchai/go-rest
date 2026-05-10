package service

import (
	"github.com/nuttchai/go-rest/internal/model"
	irepository "github.com/nuttchai/go-rest/internal/repository/interface"
)

type TUserService struct {
	Repository irepository.IUserRepository
}

func NewUserService(repository irepository.IUserRepository) *TUserService {
	return &TUserService{
		Repository: repository,
	}
}

func (s *TUserService) GetUser(id string) (*model.User, error) {
	return s.Repository.RetrieveOne(id)
}
