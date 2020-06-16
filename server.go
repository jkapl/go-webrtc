package main

import (
	// "fmt"
	// "html/template"
	"log"
	"net/http"
	// "github.com/gorilla/websocket"
)


func main() {
	fs := http.FileServer(http.Dir("./client"))
	hub := newHub()
	go hub.run()
	http.Handle("/", fs)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handleWS(w, r, hub)
	})

	log.Println("Listening on 8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}