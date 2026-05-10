//go:generate go run github.com/google/wire/cmd/wire

package di

import (
	"database/sql"
	"fmt"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"github.com/nuttchai/go-rest/internal/constant"
	ihandler "github.com/nuttchai/go-rest/internal/handler/interface"
	"github.com/nuttchai/go-rest/internal/middleware"
	"github.com/nuttchai/go-rest/internal/router"
	"github.com/nuttchai/go-rest/internal/types"
	"github.com/nuttchai/go-rest/internal/util/context"
	"github.com/nuttchai/go-rest/internal/util/env"
)

type App struct {
	Server *echo.Echo
	Config *types.TAPIConfig
}

func ProvideAPIConfig() (*types.TAPIConfig, error) {
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

	cfg := &types.TAPIConfig{
		Port: port,
		Env:  appEnv,
	}
	cfg.Db.Dsn = dbConnStr
	cfg.Db.Driver = dbDriver

	return cfg, nil
}

func ProvideDatabase(cfg *types.TAPIConfig) (*sql.DB, func(), error) {
	db, err := sql.Open(cfg.Db.Driver, cfg.Db.Dsn)
	if err != nil {
		return nil, nil, err
	}

	ctx, cancel := context.WithTimeout(constant.InitConnectionTimeout)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		_ = db.Close()
		return nil, nil, err
	}

	cleanup := func() {
		_ = db.Close()
	}

	return db, cleanup, nil
}

func ProvideServer(sampleHandler ihandler.ISampleHandler, userHandler ihandler.IUserHandler) *echo.Echo {
	e := echo.New()
	middleware.EnableCORS(e)
	router.Register(e, sampleHandler, userHandler)
	return e
}

func NewApp(server *echo.Echo, cfg *types.TAPIConfig) *App {
	return &App{
		Server: server,
		Config: cfg,
	}
}
