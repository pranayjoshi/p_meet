package chat

import (
	"p_meet/models"
	"time"
	// "github.com/fasthttp/websocket"
)

type Client struct {
	*models.Client
}

const (
	writeWait   = 10 * time.Second
	pongWait    = 60 * time.Second
	pingPeriod  = (pongWait * 9) / 10
	messageSize = 512
)

func (c *Client) readPump() {

}

func (c *Client) writePump() {

}

func PeerChatConn() {}
