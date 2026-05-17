package rest

import (
	"github.com/gofiber/fiber/v3"
)

var helthCheckResp = map[string]string{"status": "ok"}

func (s *Server) handleHealthCheck(c fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(helthCheckResp)
}
