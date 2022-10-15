package config

type APIConfig struct {
	Port string
	Env  string
	Db   struct {
		Dsn string
	}
}

type AppConfig struct {
	APIConfig APIConfig
	// Models    repositories.Models
}
