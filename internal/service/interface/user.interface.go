package iservice

import "github.com/nuttchai/go-rest/internal/model"

type IUserService interface {
	GetUser(id string) (*model.User, error)
}
