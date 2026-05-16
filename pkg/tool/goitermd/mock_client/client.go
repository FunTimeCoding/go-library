package mock_client

import (
	"github.com/funtimecoding/go-library/pkg/iterm/screen"
	"github.com/funtimecoding/go-library/pkg/iterm/session"
)

type Client struct {
	sessions []*session.Session
	screens  map[string]*screen.Screen
	nextID   int
}
