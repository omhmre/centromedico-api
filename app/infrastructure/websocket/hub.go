package websocket

import (
	"encoding/json"
	"log"
	"time"
)

// Hub maintains the set of active clients and broadcasts messages to the clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

// DeviceInfo holds the metadata of a connected user.
type DeviceInfo struct {
	User      string `json:"user"`
	Role      string `json:"role"`
	Device    string `json:"device"`
	IP        string `json:"ip"`
	Connected time.Time `json:"-"`
	Uptime    string `json:"uptime"`
}

// HubMessage is the format of JSON messages transmitted over WS.
type HubMessage struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			// We don't immediately broadcast on register because device metadata hasn't arrived yet.
			// It arrives via a "register" websocket message.

		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
				h.broadcastDeviceList()
			}

		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

// Broadcasts the list of all currently registered devices to all clients
func (h *Hub) broadcastDeviceList() {
	var devices []DeviceInfo
	now := time.Now()

	for client := range h.clients {
		if client.DeviceInfo != nil {
			uptimeDur := now.Sub(client.DeviceInfo.Connected)
			

			client.DeviceInfo.Uptime = formatDuration(uptimeDur)

			devices = append(devices, *client.DeviceInfo)
		}
	}

	msg := HubMessage{
		Type:    "devices_update",
		Payload: devices,
	}

	raw, err := json.Marshal(msg)
	if err == nil {
		for c := range h.clients {
			select {
			case c.send <- raw:
			default:
				close(c.send)
				delete(h.clients, c)
			}
		}
	} else {
		log.Printf("Error marshaling device list: %v", err)
	}
}

func formatDuration(d time.Duration) string {
	if d.Hours() > 24 {
		return "Más de 1 día"
	}
	if d.Hours() >= 1 {
		return "Hace horas"
	}
	if d.Minutes() >= 1 {
		return "Hace minutos"
	}
	return "Reciente"
}
