package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type ClientManager struct {
	clients map[*websocket.Conn]bool
	lock    sync.RWMutex
}

var clientManager = ClientManager{
	clients: make(map[*websocket.Conn]bool),
}

func HandleConnection(c *gin.Context) {

	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Failed to upgrade to websocket: %v", err)
		return
	}
	defer ws.Close()

	clientManager.addClient(ws)
	defer clientManager.removeClient(ws)

	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			break
		}
	}
}

func (cm *ClientManager) addClient(client *websocket.Conn) {
	cm.lock.Lock()
	defer cm.lock.Unlock()
	cm.clients[client] = true
}

func (cm *ClientManager) removeClient(client *websocket.Conn) {
	cm.lock.Lock()
	defer cm.lock.Unlock()
	if _, ok := cm.clients[client]; ok {
		delete(cm.clients, client)
		client.Close()
	}
}
