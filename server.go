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
	http.Handle("/", fs)
	http.HandleFunc("/ws", handleWS)
	log.Println("Listening on 8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}