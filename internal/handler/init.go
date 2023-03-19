package handler

func Init() {
	initSampleHandler()
	initUserHandler()
}

func initSampleHandler() {
	SampleHandler = &TSampleHandler{
		sampleService: service.SampleService,
		userService:   service.UserService,
	}
}

func initUserHandler() {
	UserHandler = &TUserHandler{
		userService: service.UserService,
	}
}
