package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/roger-king/go-wire-demo/pkg/repositories"
)

// StatusHandler - model to represent your status handler
type StatusHandler struct {
	Repository repositories.StatusRepositoryIface
}

// NewStatusHandler - creates and returns new instance of StatusHandler
func NewStatusHandler(r repositories.StatusRepositoryIface) *StatusHandler {
	return &StatusHandler{
		Repository: r,
	}
}

// Get - fetches the latest status of your application server
func (s *StatusHandler) Get(c *fiber.Ctx) error {
	return c.JSON(s.Repository.GetApplicationStatus())
}
