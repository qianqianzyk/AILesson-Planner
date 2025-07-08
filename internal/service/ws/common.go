package ws

import (
	"github.com/gorilla/websocket"
	"sync"
)

type WebSocketManager struct {
	connections  map[int64]map[*websocket.Conn]bool
	messageCount map[int64]int
	mu           sync.Mutex
}

func NewWebSocketManager() *WebSocketManager {
	return &WebSocketManager{
		connections:  make(map[int64]map[*websocket.Conn]bool),
		messageCount: make(map[int64]int),
	}
}

func (wsm *WebSocketManager) AddConnection(userID int64, conn *websocket.Conn) {
	wsm.mu.Lock()
	defer wsm.mu.Unlock()

	if wsm.connections[userID] == nil {
		wsm.connections[userID] = make(map[*websocket.Conn]bool)
	}
	wsm.connections[userID][conn] = true
}

func (wsm *WebSocketManager) RemoveConnection(userID int64, conn *websocket.Conn) {
	wsm.mu.Lock()
	defer wsm.mu.Unlock()

	if wsm.connections[userID] != nil {
		delete(wsm.connections[userID], conn)
		if len(wsm.connections[userID]) == 0 {
			delete(wsm.connections, userID)
		}
	}
}

func (wsm *WebSocketManager) Broadcast(userID int64, message []byte) {
	wsm.mu.Lock()
	defer wsm.mu.Unlock()

	for conn := range wsm.connections[userID] {
		err := conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			conn.Close()
			delete(wsm.connections[userID], conn)
		}
	}
}

func (wsm *WebSocketManager) IncrementMessageCount(userID int64) {
	wsm.mu.Lock()
	defer wsm.mu.Unlock()
	wsm.messageCount[userID]++
}

func (wsm *WebSocketManager) GetMessageCount(userID int64) int {
	wsm.mu.Lock()
	defer wsm.mu.Unlock()
	return wsm.messageCount[userID]
}

func (wsm *WebSocketManager) ResetMessageCount(userID int64) {
	wsm.mu.Lock()
	defer wsm.mu.Unlock()
	wsm.messageCount[userID] = 0
}
