package controller

import (
	"log"
	"octagon/dtos"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

var connections []*dtos.WebSocketConnection

const (
	MESSAGE_NEW_USER = "new_user"
	MESSAGE_CHAT     = "chat"
	MESSAGE_LEAVE    = "leave"
)

func WebSocketHandler(c *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(c) {
		return c.Next()
	}
	return fiber.ErrUpgradeRequired
}

func HandleWebSocket(conn *websocket.Conn) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR", r)
		}
	}()

	username := conn.Query("username")
	currentConn := &dtos.WebSocketConnection{Conn: conn, Username: username}
	connections = append(connections, currentConn)

	// Notify all users
	broadcastMessage(currentConn, MESSAGE_NEW_USER, "")

	for {
		var payload dtos.SocketPayload
		err := conn.ReadJSON(&payload)
		if err != nil {
			if strings.Contains(err.Error(), "websocket: close") {
				broadcastMessage(currentConn, MESSAGE_LEAVE, "")
				removeConnection(currentConn)
				return
			}
			log.Println("ERROR", err)
			continue
		}
		broadcastMessage(currentConn, MESSAGE_CHAT, payload.Message)
	}
}

func removeConnection(currentConn *dtos.WebSocketConnection) {
	var updatedConnections []*dtos.WebSocketConnection
	for _, conn := range connections {
		if conn != currentConn {
			updatedConnections = append(updatedConnections, conn)
		}
	}
	connections = updatedConnections
}

func broadcastMessage(currentConn *dtos.WebSocketConnection, kind, message string) {
	for _, conn := range connections {
		responseMessage := message

		// Jika pesan adalah MESSAGE_NEW_USER
		if kind == MESSAGE_NEW_USER {
			responseMessage = currentConn.Username + " joined chat..."
		}

		// Jika pesan adalah MESSAGE_LEAVE
		if kind == MESSAGE_LEAVE {
			responseMessage = currentConn.Username + " leaving chat..."
		}

		err := conn.Conn.WriteJSON(dtos.SocketResponse{
			From:    currentConn.Username,
			Type:    kind,
			Message: responseMessage,
		})
		if err != nil {
			log.Println("ERROR", err)
		}
	}
}
