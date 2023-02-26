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

func SetAppConfig(cfg *types.TAppConfig) {
	appConfig = cfg
}

func SetAPIConfig(cfg *types.TAPIConfig) {
	apiConfig = cfg
}

func GetAppConfig() *types.TAppConfig {
	return appConfig
}

func GetAPIConfig() *types.TAPIConfig {
	return apiConfig
}

func GetAppDB() *sql.DB {
	return appConfig.Models.DBModel.DB
}
