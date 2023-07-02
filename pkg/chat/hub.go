package chat

import "p_meet/models"

type Hub struct {
	*models.Hub
}

func NewHub() *models.Hub {
	return &models.Hub{
		Broadcast:  make(chan []byte),
		Register:   make(chan *models.Client),
		Unregister: make(chan *models.Client),
		Clients:    make(map[*models.Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
			}

		case client := <-h.Register:
			h.Clients[client] = true

		}
	}
}
