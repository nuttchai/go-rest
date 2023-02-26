package handlers

import "github.com/nuttchai/go-rest/internal/services"

func Init() {
	initSampleHandler()
	initUserHandler()
}

func initSampleHandler() {
	SampleHandler = &TSampleHandler{
		sampleService: services.SampleService,
		userService:   services.UserService,
	}
}

func initUserHandler() {
	UserHandler = &TUserHandler{
		userService: services.UserService,
	}
}
