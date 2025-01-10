package controller

import (
	"log"
	"octagon/models"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

var connections []*models.WebSocketConnection

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
	currentConn := &models.WebSocketConnection{Conn: conn, Username: username}
	connections = append(connections, currentConn)

	// Notify all users
	broadcastMessage(currentConn, MESSAGE_NEW_USER, "")

	for {
		var payload models.SocketPayload
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

func removeConnection(currentConn *models.WebSocketConnection) {
	var updatedConnections []*models.WebSocketConnection
	for _, conn := range connections {
		if conn != currentConn {
			updatedConnections = append(updatedConnections, conn)
		}
	}
	connections = updatedConnections
}

func broadcastMessage(currentConn *models.WebSocketConnection, kind, message string) {
	for _, conn := range connections {
		// if conn == currentConn {
		// 	continue
		// }
		err := conn.Conn.WriteJSON(models.SocketResponse{
			From:    currentConn.Username,
			Type:    kind,
			Message: message,
		})
		if err != nil {
			log.Println("ERROR", err)
		}
	}
}
