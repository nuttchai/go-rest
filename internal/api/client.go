package internal

import (
	"fmt"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"github.com/nuttchai/go-rest/internal/middleware"
	"github.com/nuttchai/go-rest/internal/models"
	"github.com/nuttchai/go-rest/internal/shared/config"
	"github.com/nuttchai/go-rest/internal/shared/console"
	"github.com/nuttchai/go-rest/internal/types"
)

func Client() {
	// Add the Configuration into ApiConfig
	console.App.Log("Loading App Configuration...")
	apiConfig, err := initEnv()
	if err != nil {
		console.App.Fatalf("Error Loading Root Directory (Error: %s)", err.Error())
	}
	config.SetAPIConfig(apiConfig)

	// Establish Database Connection
	console.App.Log("Connecting Database...")
	db, err := initSqlDB(config.GetAPIConfig())
	if err != nil {
		console.App.Fatalf("Database Connection Failed (Error: %s)", err.Error())
	}
	console.App.Log("Connected Database Successfully")
	defer db.Close()

	// Add the Configuration into AppConfig
	appConfig := &types.AppConfig{
		APIConfig: *config.GetAPIConfig(),
		Models:    models.Init(db),
	}
	config.SetAppConfig(appConfig)

	// Initialize Routers
	console.App.Logf("Initializing Routers...")
	e := echo.New()
	middleware.EnableCORS(e)
	initRouters(e)

	// Start Server
	console.App.Logf("Starting Server...")
	serverPort := fmt.Sprintf(":%s", config.GetAPIPort())
	if err := e.Start(serverPort); err != nil {
		console.App.Fatalf("Server Start Failed (Error: %s)", err.Error())
	}
}
