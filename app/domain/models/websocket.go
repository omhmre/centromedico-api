package models

// WebSocketMessage defines the structure for messages sent over WebSocket.
type WebSocketMessage struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data,omitempty"`
}
