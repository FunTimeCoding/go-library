package mock_listener

import "net"

func (l Listener) Addr() net.Addr {
	return nil
}
