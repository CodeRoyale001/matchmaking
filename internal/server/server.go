package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/low4ey/matchmaking/internal/config"
)

// Server represents the HTTP server.
type Server struct {
	config *config.Config
	router *mux.Router
}

// New creates a new Server instance.
func New(cfg *config.Config) *Server {
	router := mux.NewRouter()
	srv := &Server{
		config: cfg,
		router: router,
	}
	srv.routes()
	return srv
}

// Start starts the HTTP server.
func (s *Server) Start() error {
	addr := fmt.Sprintf(":%s", s.config.Port)
	log.Printf("Starting server on %s", addr)
	return http.ListenAndServe(addr, s.router)
}

// corsMiddleware is a middleware function for handling CORS.
func (s *Server) corsMiddleware(next http.Handler) http.Handler {
	allowedOrigins := []string{"http://localhost:3000", "https://serene-fortress-91389-77d1fb95872a.herokuapp.com", "https://coderoyale.vercel.app"}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		// Check if origin is present and allowed
		if origin != "" {
			for _, allowedOrigin := range allowedOrigins {
				if origin == allowedOrigin {
					w.Header().Set("Access-Control-Allow-Origin", origin)
					break
				}
			}
		}

		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, origin, x-requested-with")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
