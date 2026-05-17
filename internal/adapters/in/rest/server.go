package rest

import (
	"github.com/bignuttawuts/go-async-job-application-api/config"
	"github.com/bignuttawuts/go-async-job-application-api/internal/usecases"
	"github.com/gofiber/fiber/v3"
)

type Server struct {
	app    *fiber.App
	config config.Config
	uc     *UseCases
}

type UseCases struct {
	CreateApplicationUsecase *usecases.CreateApplicationUsecase
}

func NewServer(uc *UseCases, config config.Config) *Server {
	s := &Server{uc: uc, config: config}
	s.app = fiber.New(s.buildConfig())
	return s
}

func (s *Server) buildConfig() fiber.Config {
	return fiber.Config{
		// DisableStartupMessage: true,
		Immutable: true,
		BodyLimit: 10 * 1024 * 1024, // 10MB
		// Prefork: //s.config.Prefork,
		ReadTimeout:  s.config.ReadTimeout,
		WriteTimeout: s.config.WriteTimeout,
	}
}

func (s *Server) ListenAndServe() error {
	s.routes()
	return s.app.Listen(s.config.Port)
}
