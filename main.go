package main

import (
	"log"
	"net/http"
)

func main() {
	hub := newHub()
	go hub.run()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		// recieve request at ws://localhost/ws
		// now we call serveWs to upgrade http to ws
		serveWs(hub, w, r)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
