package irepository

import "github.com/nuttchai/go-rest/internal/model"

type IUserRepository interface {
	RetrieveOne(id string) (*model.User, error)
}
