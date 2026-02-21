package connection

import (
	"context"
	"github.com/gorilla/websocket"
	"sync"
)

type Connection struct {
	sync.RWMutex
	host           string
	token          string
	connection     *websocket.Conn
	subscribers    map[uint64]Subscriber
	lastIdentifier uint64
	context        context.Context
	cancel         context.CancelFunc
}
