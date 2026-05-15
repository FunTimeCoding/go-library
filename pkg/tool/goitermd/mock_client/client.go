package mock_client

import "github.com/funtimecoding/go-library/pkg/iterm/session"

type Client struct {
	sessions []session.Session
	screens  map[string]session.Screen
	nextID   int
}
