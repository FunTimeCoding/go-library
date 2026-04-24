package firefox

import (
	"github.com/coder/websocket"
	"sync"
	"sync/atomic"
)

type Client struct {
	address    string
	connection *websocket.Conn
	mutex      sync.Mutex
	pending    map[int]chan response
	identifier atomic.Int64
}
