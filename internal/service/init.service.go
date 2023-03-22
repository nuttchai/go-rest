package service

import "github.com/nuttchai/go-rest/internal/repository"

func Init() {
	initSampleService(&TSampleService{
		Repository: repository.SampleRepository,
	})

	initUserService(&TUserService{
		Repository: repository.UserRepository,
	})
}
