package chat

import "p_meet/models"

func NewHub() *models.Hub {
	return &models.Hub{
		Broadcast:  make(chan []byte),
		Register:   make(chan *models.Client),
		Unregister: make(chan *models.Client),
		Clients:    make(map[*models.Client]bool),
	}
}

func (h *models.Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
			}

		}
	}
}
