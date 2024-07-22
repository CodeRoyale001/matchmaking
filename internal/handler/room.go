package handler

type BroadcastMessage struct {
	matchId string
	message []byte
}

type Hub struct {
	clients    map[*Client]bool
	broadcast  chan BroadcastMessage
	register   chan *Client
	unregister chan *Client
	matchRooms map[string]map[*Client]bool
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan BroadcastMessage),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
		matchRooms: make(map[string]map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			if _, ok := h.matchRooms[client.matchId]; !ok {
				h.matchRooms[client.matchId] = make(map[*Client]bool)
			}
			h.matchRooms[client.matchId][client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
			if room, ok := h.matchRooms[client.matchId]; ok {
				if _, ok := room[client]; ok {
					delete(room, client)
					if len(room) == 0 {
						delete(h.matchRooms, client.matchId)
					}
				}
			}
		case broadcastDeatil := <-h.broadcast:
			if room, ok := h.matchRooms[broadcastDeatil.matchId]; ok {
				for client := range room {
					select {
					case client.send <- broadcastDeatil.message:
					default:
						close(client.send)
						delete(h.clients, client)
						delete(room, client)
					}
				}
			}
		}
	}
}
