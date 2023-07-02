package webrtc

import (
	"log"
	"p_meet/models"
	"p_meet/pkg/webrtc"
	"sync"

	"github.com/gofiber/websocket/v2"
)

func RoomConn(c *websocket.Conn, p *models.Peers) {
	var config webrtc.Configuration

	peerConnection, err := webrtc.NewPeerConnection(config)
	if err != nil {
		log.Print(err)
		return
	}

	newPeer := PeerConnectionState{
		PeerConnection: peerConnection,
		WebSocket:      &ThreadSafeWriter,
		Conn:           c,
		Mutex:          sync.Mutex{},
	}
}
