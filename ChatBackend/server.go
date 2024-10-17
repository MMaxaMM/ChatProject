package chat

import (
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(address string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:    address,
		Handler: handler,
	}

	return s.httpServer.ListenAndServe()
}
