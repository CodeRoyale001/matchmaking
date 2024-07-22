package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/low4ey/matchmaking/package/db"
	"github.com/low4ey/matchmaking/package/models"
	"github.com/low4ey/matchmaking/package/utils"
)

var redisClient *redis.Client

func init() {
	redisClient = db.Connect()
}

// NewMatch creates a new Match instance with default values.
func NewMatch(matchID string) *models.Match {
	return &models.Match{
		MatchID:     matchID,
		MaxPlayers:  2, // Default value for max_players
		PlayerAID:   []string{},
		IsAvailable: true,
	}
}

func SearchMatch(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	// Decode the request body
	playerID := r.URL.Query().Get("playerID")
	if playerID == "" {
		http.Error(w, "Missing playerID parameter", http.StatusBadRequest)
		return
	}
	for {
		currMatch, err := redisClient.ZRevRangeByScoreWithScores(ctx, "Match:1v1", &redis.ZRangeBy{
			Min:    "-inf",
			Max:    "inf",
			Offset: 0,
			Count:  1,
		}).Result()
		if err != nil {
			utils.SendErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		if len(currMatch) == 0 {
			break
		} else {
			var matchData models.Match
			for _, match := range currMatch {
				json.Unmarshal([]byte(match.Member.(string)), &matchData)
				redisClient.ZRem(ctx, "Match:1v1", match.Member)
				if matchData.PlayerCount == matchData.MaxPlayers {
					matchData.IsAvailable = false
				} else if matchData.IsAvailable {
					matchData.PlayerCount++
					matchData.PlayerAID = append(matchData.PlayerAID, playerID)
					matchDataBytes, _ := json.Marshal(matchData)
					redisClient.ZAdd(ctx, "Match:1v1", &redis.Z{Score: float64(matchData.PlayerCount), Member: string(matchDataBytes)})
					utils.SendJSONResponse(w, http.StatusOK, "Match Found", matchData)
					return
				}
			}
			break
		}
	}
	fmt.Println(playerID, "searching for match")
	newMatchId := uuid.New().String()
	newMatch := NewMatch(newMatchId)
	newMatch.PlayerCount++
	newMatch.PlayerAID = append(newMatch.PlayerAID, playerID)
	matchData, _ := json.Marshal(newMatch)
	redisClient.ZAdd(ctx, "Match:1v1", &redis.Z{Score: float64(newMatch.PlayerCount), Member: string(matchData)})
	utils.SendJSONResponse(w, http.StatusOK, "Match found or created", newMatch)
}
