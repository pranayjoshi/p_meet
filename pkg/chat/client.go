package chat

import (
	"p_meet/models"
	"time"
	// "github.com/fasthttp/websocket"
)

const (
	writeWait   = 10 * time.Second
	pongWait    = 60 * time.Second
	pingPeriod  = (pongWait * 9) / 10
	messageSize = 512
)

func (c *models.Client) readPump() {

}

func (c *models.Client) writePump() {

}

func PeerChatConn() {}
