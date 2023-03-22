package handler

import (
	"github.com/nuttchai/go-rest/internal/service"
)

func Init() {
	initSampleHandler(&TSampleHandler{
		SampleService: service.SampleService,
		UserService:   service.UserService,
	})

	initUserHandler(&TUserHandler{
		UserService: service.UserService,
	})
}

func initSampleHandler(sampleHandler *TSampleHandler) {
	SampleHandler = sampleHandler
}

func initUserHandler(userHandler *TUserHandler) {
	UserHandler = userHandler
}
