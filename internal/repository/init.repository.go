package repository

import "github.com/nuttchai/go-rest/internal/shared/config"

func Init() {
	initSampleRepository(&TSampleRepository{
		DB: config.GetAppDB(),
	})

	initUserRepository(&TUserRepository{
		DB: config.GetAppDB(),
	})
}
