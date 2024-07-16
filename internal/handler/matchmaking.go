package handler

import (
	"net/http"

	"github.com/low4ey/matchmaking/package/utils"
)

// UserHandler handles requests to the /users endpoint.
func Hello(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{"hello": "world"}
	utils.SendJSONResponse(w, http.StatusOK, "User retrieved successfully", data)
}
