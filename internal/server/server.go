package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/low4ey/matchmaking/internal/config"
	"github.com/low4ey/matchmaking/internal/handler"
)

// Server represents the HTTP server.
type Server struct {
	config *config.Config
	mux    *http.ServeMux
}

// New creates a new Server instance.
func New(cfg *config.Config) *Server {
	mux := http.NewServeMux()
	srv := &Server{
		config: cfg,
		mux:    mux,
	}
	srv.routes()
	return srv
}

// Start starts the HTTP server.
func (s *Server) Start() error {
	addr := fmt.Sprintf(":%s", s.config.Port)
	log.Printf("Starting server on %s", addr)
	return http.ListenAndServe(addr, s.mux)
}

// routes registers the HTTP routes.
func (s *Server) routes() {
	s.mux.HandleFunc("/", handler.Hello)
	// s.mux.HandleFunc("/products", handlers.ProductHandler)
}
