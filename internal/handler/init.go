package handler

import "github.com/nuttchai/go-rest/internal/service"

func Init() {
	initSampleHandler()
	initUserHandler()
}

func initSampleHandler() {
	SampleHandler = &TSampleHandler{
		SampleService: service.SampleService,
		UserService:   service.UserService,
	}
}

func initUserHandler() {
	UserHandler = &TUserHandler{
		UserService: service.UserService,
	}
}
