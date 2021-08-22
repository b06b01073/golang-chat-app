package main

type Hub struct {
	// Registered clients.
	// This map will be use to push the message to all clients
	clients map[*Client]bool

	// this channel is use to store the message from the client that will be broadcast.
	broadcast chan []byte

	//  registerd client channel is for client that build the connection.
	register chan *Client

	// unregister client channel is for client that is going to be disconnected.
	unregister chan *Client
}

// newHub return a Hub reference with zero value at all fields
func newHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

// run a hub
func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			// this case means there is a waiting client which is ready to register
			// we mark this client in the clients map as true, so the upcomming broadcasts will be send to this client
			h.clients[client] = true
		case client := <-h.unregister:
			// this case means there is a waiting client which is ready to unregister
			// we delete this client from the clients map, and close the send channel, which make sure there are no more
			// write to the channel
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			// this case means there are a waiting message which is ready to be broadcasted
			for client := range h.clients {
				// for every clients in the clients map, we send the message to them
				select {
				case client.send <- message:
					// this case means the send channel of the client is ready(not blocked) to recieve message
				default:
					// if the send channel is block, then we close the channel
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
