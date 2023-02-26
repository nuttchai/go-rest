package repositories

import "github.com/nuttchai/go-rest/internal/shared/config"

func Init() {
	initSampleRepository()
	initUserRepository()
}

func initSampleRepository() {
	SampleRepository = &TSampleRepository{
		DB: config.GetAppDB(),
	}
}

func initUserRepository() {
	UserRepository = &TUserRepository{
		DB: config.GetAppDB(),
	}
}
