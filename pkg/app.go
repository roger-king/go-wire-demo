package pkg

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/google/wire"
	"github.com/roger-king/go-wire-demo/pkg/config"
	"github.com/roger-king/go-wire-demo/pkg/handlers"
	"github.com/roger-king/go-wire-demo/pkg/repositories"
	log "github.com/sirupsen/logrus"
)

// AppSet -
var AppSet = wire.NewSet(NewApp, repositories.Set, handlers.Set, config.Set)

// NewApp -
func NewApp(status *handlers.StatusHandler) *fiber.App {
	app := fiber.New()

	app.Use(cors.New())

	api := app.Group("/api")
	// Server status endpoint - sanity check that the server is running
	statusGroup := api.Group("/status")
	statusGroup.Get("/", status.Get)

	log.Info("Application is running...")
	return app
}
