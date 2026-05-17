package rest

import (
	"fmt"

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
		// Prefork: //s.config.Server.Prefork,
		ReadTimeout:  s.config.Server.ReadTimeout,
		WriteTimeout: s.config.Server.WriteTimeout,
	}
}

func (s *Server) ListenAndServe() error {
	s.routes()
	return s.app.Listen(fmt.Sprintf(":%s", s.config.Server.Port))
}
