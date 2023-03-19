package types

import "github.com/nuttchai/go-rest/internal/model"

type TAPIConfig struct {
	Port string
	Env  string
	Db   struct {
		Dsn    string
		Driver string
	}
}

type TAppConfig struct {
	APIConfig TAPIConfig
	Models    model.Models
}
