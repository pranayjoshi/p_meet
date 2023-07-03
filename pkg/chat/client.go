package chat

import (
	"p_meet/models"
	"time"

	// "github.com/gofiber/websocket/v2"
	"github.com/fasthttp/websocket"
)

type Client struct {
	*models.Client
}

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.FastHTTPUpgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
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
