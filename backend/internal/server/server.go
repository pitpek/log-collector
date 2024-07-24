package server

import (
	"context"
	"net/http"
	"strconv"
	"time"
)

// Server представляет собой HTTP-сервер.
type Server struct {
	httpServer *http.Server
}

// Run запускает HTTP-сервер на указанном порту с заданным обработчиком.
func (s *Server) Run(port int, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           "0.0.0.0:" + strconv.Itoa(port),
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return s.httpServer.ListenAndServe()
}

// Shutdown останавливает HTTP-сервер.
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
