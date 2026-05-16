package mock_client

import "github.com/funtimecoding/go-library/pkg/sublime/view"

type Client struct {
	views  []*view.View
	nextID int
}
