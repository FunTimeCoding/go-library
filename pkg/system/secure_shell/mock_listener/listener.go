package mock_listener

import (
	"github.com/funtimecoding/go-library/pkg/system/secure_shell/mock_connection"
	"net"
)

type Listener struct{}

func (l Listener) Accept() (net.Conn, error) {
	return mock_connection.New(), nil
}

func (l Listener) Close() error {
	return nil
}

func (l Listener) Addr() net.Addr {
	return nil
}
