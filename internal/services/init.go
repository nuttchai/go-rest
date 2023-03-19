package services

import repository "github.com/nuttchai/go-rest/internal/repository"

func Init() {
	initUserService()
	initSampleService()
}

func initSampleService() {
	SampleService = &TSampleService{
		repository: repository.SampleRepository,
	}
}

func initUserService() {
	UserService = &TUserService{
		repository: repository.UserRepository,
	}
}
