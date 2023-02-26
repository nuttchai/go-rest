package config

import (
	"database/sql"

	"github.com/nuttchai/go-rest/internal/types"
)

var appConfig *types.AppConfig
var apiConfig *types.APIConfig

func init() {
	appConfig = &types.AppConfig{}
	apiConfig = &types.APIConfig{}
}

// AppConfig is the global configuration for the application
func SetAppConfig(cfg *types.AppConfig) {
	appConfig = cfg
}

func GetAppConfig() *types.AppConfig {
	return appConfig
}

func GetAppDB() *sql.DB {
	return appConfig.Models.DBModel.DB
}

// APIConfig is the global configuration for the API
func SetAPIConfig(cfg *types.APIConfig) {
	apiConfig = cfg
}

func GetAPIConfig() *types.APIConfig {
	return apiConfig
}

func GetAPIPort() string {
	return apiConfig.Port
}
