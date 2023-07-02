package models

import (
	"p_meet/pkg/chat"
)

type Room struct {
	Peers *Peers
	Hub   *chat.Hub
}
