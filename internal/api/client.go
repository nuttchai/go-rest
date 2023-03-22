package internal

import (
	"fmt"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"github.com/nuttchai/go-rest/internal/middleware"
	"github.com/nuttchai/go-rest/internal/model"
	"github.com/nuttchai/go-rest/internal/shared/config"
	"github.com/nuttchai/go-rest/internal/shared/console"
	"github.com/nuttchai/go-rest/internal/types"
)

func init() {

}
func Client() {
	// Add the Configuration into ApiConfig
	console.App.Log("Loading App Configuration...")
	cfg, err := initEnv()
	if err != nil {
		console.App.Fatalf("Error Loading Root Directory (Error: %s)", err.Error())
	}
	config.SetAPIConfig(cfg)

	// Establish Database Connection
	console.App.Log("Connecting Database...")
	apiConfig := config.GetAPIConfig()
	db, err := initSqlDB(apiConfig)
	if err != nil {
		console.App.Fatalf("Database Connection Failed (Error: %s)", err.Error())
	}
	console.App.Log("Connected Database Successfully")
	defer db.Close()

	// Add the Configuration into AppConfig
	appConfig := &types.TAppConfig{
		APIConfig: *apiConfig,
		Models:    model.Init(db),
	}
	config.SetAppConfig(appConfig)

	// Initialize App
	console.App.Logf("Initializing the Application...")
	e := echo.New()
	middleware.EnableCORS(e)
	initApp(e)

	// Start Server
	console.App.Logf("Starting Server...")
	serverPort := fmt.Sprintf(":%s", apiConfig.Port)
	if err := e.Start(serverPort); err != nil {
		console.App.Fatalf("Server Start Failed (Error: %s)", err.Error())
	}
}
