package mock_listener

import (
	"github.com/funtimecoding/go-library/pkg/system/secure_shell/mock_connection"
	"net"
)

func (l Listener) Accept() (net.Conn, error) {
	return mock_connection.New(), nil
}
