package types

import (
	"github.com/nuttchai/go-rest/internal/models"
)

type APIConfig struct {
	Port string
	Env  string
	Db   struct {
		Dsn    string
		Driver string
	}
}

type AppConfig struct {
	APIConfig APIConfig
	Models    models.Models
}
