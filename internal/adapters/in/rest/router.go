package rest

func (s *Server) routes() {
	s.app.Get("/health", s.handleHealthCheck)

	root := s.app.Group("/api/v1/jobs/:jobID/applications")
	root.Post("/", s.handleCreateApplication)
}
