package main

type Hub struct {
	clients map[*Client]bool
	broadcast chan []byte
	register chan *Client
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		clients: make(map[*Client]bool)
		broadcast: make(chan []byte)
		register: make(chan *Client)
		unregister: make(chan *Client)
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