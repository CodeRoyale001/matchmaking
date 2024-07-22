package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// ErrorResponse represents the structure of an error response.
type ErrorResponse struct {
	Message string `json:"message"`
}

// JSONResponse represents the structure of a standard JSON response.
type JSONResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// SendJSONResponse sends a JSON response with the given status code, message, and data.
func SendJSONResponse(w http.ResponseWriter, status int, message string, data interface{}) {
	response := JSONResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// SendErrorResponse sends a JSON error response with the given status code and message.
func SendErrorResponse(w http.ResponseWriter, status int, message string) {
	response := ErrorResponse{
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Read WebSocket messages
func ReadMessage(conn *websocket.Conn) (message []byte, err error) {
	_, msg, err := conn.ReadMessage()
	if err != nil {
		log.Println("Error reading message:", err)
		return msg, err
	}
	log.Printf("Received message: %s", msg)
	return msg, err
}

// Write WebSocket messages
func WriteMessage(conn *websocket.Conn, msg []byte) {
	err := conn.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		log.Println("Error writing message:", err)
		return
	}
	log.Printf("Sent message: %s", msg)
}
