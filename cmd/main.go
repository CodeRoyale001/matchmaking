package main

import (
	"log"

	"github.com/low4ey/matchmaking/internal/config"
	"github.com/low4ey/matchmaking/internal/server"
	"github.com/low4ey/matchmaking/package/db"
)

func main() {
	db.Connect()
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	srv := server.New(cfg)
	if err := srv.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
