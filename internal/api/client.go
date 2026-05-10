package internal

import (
	"fmt"

	"github.com/nuttchai/go-rest/internal/di"
	"github.com/nuttchai/go-rest/internal/shared/console"
)

func init() {

}
func Client() {
	console.App.Log("Initializing the application dependencies...")
	app, cleanup, err := di.InitializeApp()
	if err != nil {
		console.App.Fatalf("Failed to initialize application dependencies (Error: %s)", err.Error())
	}
	defer cleanup()

	console.App.Log("Application dependencies initialized")
	console.App.Log("Starting Server...")
	serverPort := fmt.Sprintf(":%s", app.Config.Port)
	if err := app.Server.Start(serverPort); err != nil {
		console.App.Fatalf("Server Start Failed (Error: %s)", err.Error())
	}
}
