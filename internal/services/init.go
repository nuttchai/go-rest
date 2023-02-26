package services

import "github.com/nuttchai/go-rest/internal/repositories"

func Init() {
	initUserService()
	initSampleService()
}

func initSampleService() {
	SampleService = &TSampleService{
		repository: repositories.SampleRepository,
	}
}

func initUserService() {
	UserService = &TUserService{
		repository: repositories.UserRepository,
	}
}
