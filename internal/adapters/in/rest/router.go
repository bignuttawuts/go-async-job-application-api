package rest

func (s *Server) routes() {
	s.app.Get("/health", s.handleHealthCheck)
}
