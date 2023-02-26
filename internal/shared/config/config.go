package config

import (
	"database/sql"

	"github.com/nuttchai/go-rest/internal/types"
)

var appConfig *types.TAppConfig
var apiConfig *types.TAPIConfig

func init() {
	appConfig = &types.TAppConfig{}
	apiConfig = &types.TAPIConfig{}
}

// AppConfig is the global configuration for the application
func SetAppConfig(cfg *types.TAppConfig) {
	appConfig = cfg
}

func GetAppConfig() *types.TAppConfig {
	return appConfig
}

func GetAppDB() *sql.DB {
	return appConfig.Models.DBModel.DB
}

// APIConfig is the global configuration for the API
func SetAPIConfig(cfg *types.TAPIConfig) {
	apiConfig = cfg
}

func GetAPIConfig() *types.TAPIConfig {
	return apiConfig
}

func GetAPIPort() string {
	return apiConfig.Port
}
