package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/low4ey/matchmaking/package/utils"
)

// PlayerRequest represents the request body containing player_id.
type PlayerRequest struct {
	PlayerID string `json:"player_id"`
}

type Match struct {
	MatchID     string
	MaxPlayers  int
	PlayerCount int
	PlayerAID   []string
	IsAvailable bool
}

// NewMatch creates a new Match instance with default values.
func NewMatch(matchID string) *Match {
	return &Match{
		MatchID:     matchID,
		MaxPlayers:  2, // Default value for max_players
		PlayerAID:   []string{},
		IsAvailable: true,
	}
}

var availableMatch []Match

// UserHandler handles requests to the /users endpoint.
func Hello(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{"hello": "world"}
	utils.SendJSONResponse(w, http.StatusOK, "User retrieved successfully", data)
}

func SearchMatch(w http.ResponseWriter, r *http.Request) {
	// Decode the request body
	var req PlayerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	playerID := req.PlayerID
	fmt.Println(playerID, "searching for match")
	fmt.Println(availableMatch)
	for index := range availableMatch {
		if availableMatch[index].IsAvailable {
			currMatch := &availableMatch[index]
			currMatch.PlayerCount++
			currMatch.IsAvailable = currMatch.PlayerCount < currMatch.MaxPlayers
			currMatch.PlayerAID = append(currMatch.PlayerAID, playerID)
			utils.SendJSONResponse(w, http.StatusOK, "Match found", currMatch)
			return
		}
	}
	newMatchId := uuid.New().String()
	newMatch := NewMatch(newMatchId)
	newMatch.PlayerCount++
	newMatch.PlayerAID = append(newMatch.PlayerAID, playerID)
	availableMatch = append(availableMatch, *newMatch)
	utils.SendJSONResponse(w, http.StatusOK, "Match created", newMatch)
}
