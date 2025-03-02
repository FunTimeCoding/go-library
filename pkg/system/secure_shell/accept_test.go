package secure_shell

import (
	"net"
	"testing"
	"time"
)

type ConnMock struct{}

func (c ConnMock) Read(_ []byte) (n int, err error) {
	return 0, nil
}

func (c ConnMock) Write(_ []byte) (n int, err error) {
	return 0, nil
}

func (c ConnMock) Close() error {
	return nil
}

func (c ConnMock) LocalAddr() net.Addr {
	return nil
}

func (c ConnMock) RemoteAddr() net.Addr {
	return nil
}

func (c ConnMock) SetDeadline(_ time.Time) error {
	return nil
}

func (c ConnMock) SetReadDeadline(_ time.Time) error {
	return nil
}

func (c ConnMock) SetWriteDeadline(_ time.Time) error {
	return nil
}

type ListenerMock struct {
}

func (l ListenerMock) Accept() (net.Conn, error) {
	return &ConnMock{}, nil
}

func (l ListenerMock) Close() error {
	return nil
}

func (l ListenerMock) Addr() net.Addr {
	return nil
}

func TestAccept(t *testing.T) {
	Accept(&ListenerMock{})
}
