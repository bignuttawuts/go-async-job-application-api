package rest

import (
	"github.com/gofiber/fiber/v3"
)

var helthCheckResp = map[string]string{"status": "ok"}

func (s *Server) handleHealthCheck(c fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(helthCheckResp)
}

func (s *Server) handleCreateApplication(c fiber.Ctx) error {
	if err := s.uc.CreateApplicationUsecase.Execute(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]string{"error": err.Error()})
	}
	return c.Status(fiber.StatusAccepted).JSON(map[string]string{"status": "ok", "message": "Application received and is being processed."})
}
