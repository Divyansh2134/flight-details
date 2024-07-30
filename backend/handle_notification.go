package main

import "github.com/gorilla/websocket"

func NotifyClients(message string) {
	clientManager.lock.RLock()
	defer clientManager.lock.RUnlock()

	for client := range clientManager.clients {
		err := client.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			Log.Error("Error writing message to client: %v", err)
			client.Close()
			clientManager.removeClient(client)
		}
	}
}
