package models

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
