package main

type Hub struct {
	clients map[*websocket.Conn]bool
	broadcast chan []byte
	register chan *websocket.Conn
	unregister chan *websocket.Conn
}

func newHub() *Hub {
	return &Hub{
		clients: make(map[*websocket.Conn]bool)
		broadcast: make(chan []byte)
		register: make(chan *websocket.Conn)
		unregister: make(chan *websocket.Conn)
	}
}

func (h *Hub) run() {
	for {
		select {
			case client := <-h.register
				h.clients[client] = true
			case client := <-h.unregister:
				// if client is in clients map delete the key and close the channel (delete and close are built-in functions)
				if _, ok := h.clients[client]; ok {
					delete(h.clients, client)
					close(client.send)
				}
			case message := <-h.broadcast:
				for client := range h.clients {
					client.send <- message
				}
		}
	}
}