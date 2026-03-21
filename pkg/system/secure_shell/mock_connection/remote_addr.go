package mock_connection

import "net"

func (c Connection) RemoteAddr() net.Addr {
	return nil
}
