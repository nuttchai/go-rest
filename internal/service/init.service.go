package service

import repository "github.com/nuttchai/go-rest/internal/repository"

func Init() {
	initUserService()
	initSampleService()
}

func initSampleService() {
	SampleService = &TSampleService{
		Repository: repository.SampleRepository,
	}
}

func initUserService() {
	UserService = &TUserService{
		Repository: repository.UserRepository,
	}
}
