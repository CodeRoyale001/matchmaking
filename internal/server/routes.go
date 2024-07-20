package server

import (
	"net/http"

	"github.com/low4ey/matchmaking/internal/handler"
)

// routes registers the HTTP routes.
func (s *Server) routes() {
	s.router.Handle("/ping", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Pong"))
	}))
	s.router.Handle("/search", s.corsMiddleware(http.HandlerFunc(handler.SearchMatch)))
	// s.mux.Handle("/products", s.corsMiddleware(http.HandlerFunc(handlers.ProductHandler)))
}
