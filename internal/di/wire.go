//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/nuttchai/go-rest/internal/handler"
	ihandler "github.com/nuttchai/go-rest/internal/handler/interface"
	"github.com/nuttchai/go-rest/internal/repository"
	irepository "github.com/nuttchai/go-rest/internal/repository/interface"
	"github.com/nuttchai/go-rest/internal/service"
	iservice "github.com/nuttchai/go-rest/internal/service/interface"
)

var repositorySet = wire.NewSet(
	repository.NewSampleRepository,
	wire.Bind(new(irepository.ISampleRepository), new(*repository.TSampleRepository)),
	repository.NewUserRepository,
	wire.Bind(new(irepository.IUserRepository), new(*repository.TUserRepository)),
)

var serviceSet = wire.NewSet(
	service.NewSampleService,
	wire.Bind(new(iservice.ISampleService), new(*service.TSampleService)),
	service.NewUserService,
	wire.Bind(new(iservice.IUserService), new(*service.TUserService)),
)

var handlerSet = wire.NewSet(
	handler.NewSampleHandler,
	wire.Bind(new(ihandler.ISampleHandler), new(*handler.TSampleHandler)),
	handler.NewUserHandler,
	wire.Bind(new(ihandler.IUserHandler), new(*handler.TUserHandler)),
)

//go:generate go run github.com/google/wire/cmd/wire
func InitializeApp() (*App, func(), error) {
	wire.Build(
		ProvideAPIConfig,
		ProvideDatabase,
		ProvideServer,
		NewApp,
		repositorySet,
		serviceSet,
		handlerSet,
	)
	return nil, nil, nil
}
