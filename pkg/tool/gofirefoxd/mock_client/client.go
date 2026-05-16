package mock_client

import "github.com/funtimecoding/go-library/pkg/firefox/tab"

type Client struct {
	tabs    []*tab.Tab
	groups  map[int]*group
	nextID  int
	groupID int
}
