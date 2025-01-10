package models

import "github.com/gofiber/websocket/v2"

type WebSocketConnection struct {
	Conn     *websocket.Conn
	Username string
}

type SocketPayload struct {
	Message string `json:"message"`
}

type SocketResponse struct {
	From    string `json:"from"`
	Type    string `json:"type"`
	Message string `json:"message"`
}
