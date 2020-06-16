package main

import (
	"encoding/json"
	"io"
	"net/http"
	"github.com/gorilla/websocket"
)


func handleWS(w http.ResponseWriter, r *http.Request, hub *Hub) {
    conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}


	hub.register <- conn
	

}
