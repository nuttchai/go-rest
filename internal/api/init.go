package internal

import (
	"database/sql"
	"flag"
	"fmt"

	"github.com/labstack/echo"
	"github.com/nuttchai/go-rest/internal/constant"
	"github.com/nuttchai/go-rest/internal/handler"
	"github.com/nuttchai/go-rest/internal/repository"
	"github.com/nuttchai/go-rest/internal/router"
	"github.com/nuttchai/go-rest/internal/service"
	"github.com/nuttchai/go-rest/internal/types"
	"github.com/nuttchai/go-rest/internal/util/context"
	"github.com/nuttchai/go-rest/internal/util/env"
)

func initEnv() (*types.TAPIConfig, error) {
	apiConfig := &types.TAPIConfig{}

	// Load Environment Variables
	appEnv := env.GetEnv("APP_ENV", "development")
	envDefaultDir, err := env.GetDefaultEnvDir(appEnv)
	if err != nil {
		return nil, err
	}

	envDir := env.GetEnv("ENV_PATH", envDefaultDir)
	env.LoadEnv(envDir)

	dbType := env.GetEnv("DB_TYPE", "postgres")
	dbUser := env.GetEnv("APP_DB_USER", "postgres")
	dbPass := env.GetEnv("APP_DB_PASS", "postgres")
	dbHost := env.GetEnv("DB_HOST", "localhost")
	dbPort := env.GetEnv("DB_PORT", "5432")
	dbName := env.GetEnv("APP_DB_NAME", "test")
	dbDriver := env.GetEnv("DB_DRIVER", "postgres")
	port := env.GetEnv("APP_PORT", "8000")
	dbConnStr := fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s?sslmode=disable",
		dbType,
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName,
	)

	// Store ENV Variables to ApiConfig
	flag.StringVar(&apiConfig.Env, "env", appEnv, "Application Environment")
	flag.StringVar(&apiConfig.Port, "port", port, "Server Listening Port")
	flag.StringVar(&apiConfig.Db.Dsn, "dsn", dbConnStr, "Data Source Name")
	flag.StringVar(&apiConfig.Db.Driver, "driver", dbDriver, "Database Driver")
	flag.Parse()

	return apiConfig, nil
}

func initSqlDB(cfg *types.TAPIConfig) (*sql.DB, error) {
	db, err := sql.Open(cfg.Db.Driver, cfg.Db.Dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(constant.InitConnectionTimeout)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func initApp(e *echo.Echo) {
	repository.Init()
	service.Init()
	handler.Init()
	router.Init(e)
}
