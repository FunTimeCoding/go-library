package mock_connection

import (
	"net"
	"time"
)

type Connection struct{}

func (c Connection) Read(_ []byte) (n int, err error) {
	return 0, nil
}

func (c Connection) Write(_ []byte) (n int, err error) {
	return 0, nil
}

func (c Connection) Close() error {
	return nil
}

func (c Connection) LocalAddr() net.Addr {
	return nil
}

func (c Connection) RemoteAddr() net.Addr {
	return nil
}

func (c Connection) SetDeadline(_ time.Time) error {
	return nil
}

func (c Connection) SetReadDeadline(_ time.Time) error {
	return nil
}

func (c Connection) SetWriteDeadline(_ time.Time) error {
	return nil
}
