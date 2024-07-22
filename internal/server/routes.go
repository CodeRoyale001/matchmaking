package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/low4ey/matchmaking/internal/handler"
)

// routes registers the HTTP routes.
func (s *Server) routes() {
	hub := handler.NewHub()
	go hub.Run()
	s.router.HandleFunc("/ws/{MatchId}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		matchId := vars["MatchId"]
		handler.ServeWs(hub, matchId, w, r)
	}).Methods("GET")
	s.router.Handle("/ping", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Pong"))
	}))
	s.router.Handle("/search", s.corsMiddleware(http.HandlerFunc(handler.SearchMatch)))
}
