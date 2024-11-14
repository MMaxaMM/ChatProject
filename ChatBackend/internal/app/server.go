package app

import (
	"chat/internal/config"
	"net/http"
)

type Server struct {
	httpServer *http.Server
	cfg        config.HTTPServer
}

func NewServer(cfg config.HTTPServer) *Server {
	return &Server{cfg: cfg}
}

func (s *Server) Run(handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:    s.cfg.Address,
		Handler: handler,
	}

	return s.httpServer.ListenAndServe()
}
