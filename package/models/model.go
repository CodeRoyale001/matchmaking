package models

import "github.com/gorilla/websocket"

// User represents a user in the system.
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// Match represents a match in the system.
type Match struct {
	MatchID     string   `json:"match_id"`
	MaxPlayers  int      `json:"max_players"`
	PlayerCount int      `json:"player_count"`
	PlayerAID   []string `json:"player_a_id"`
	IsAvailable bool     `json:"is_available"`
}

// Player Request represents the request body containing player_id
type PlayerRequest struct {
	PlayerID string `json:"player_id"`
}

// Room represents a room in the system.
type room struct{}

// Player represents a player in the system.
type Player struct {
	socket  *websocket.Conn
	receive chan []byte
	room    *room
}
type WsRequest struct {
	Action  string `json:"action"`
	Match   string `json:"match"`
	Content string `json:"content"`
}
