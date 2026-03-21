package mock_connection

import "net"

func (c Connection) LocalAddr() net.Addr {
	return nil
}
