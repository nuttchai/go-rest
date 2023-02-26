package types

import (
	"github.com/nuttchai/go-rest/internal/models"
)

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
	Models    models.Models
}
