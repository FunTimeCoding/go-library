package firefox

import (
	"encoding/json"
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

type request struct {
	Method     string `json:"method"`
	Parameters any    `json:"params"`
	Identifier int    `json:"id"`
}

type response struct {
	Result     json.RawMessage `json:"result,omitempty"`
	Error      string          `json:"error,omitempty"`
	Identifier int             `json:"id"`
}
