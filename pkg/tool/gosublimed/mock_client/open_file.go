package mock_client

import "github.com/funtimecoding/go-library/pkg/sublime/view"

func (c *Client) OpenFile(path string) (view.View, error) {
	v := view.View{
		Identifier: c.nextID,
		FilePath:   path,
	}
	c.nextID++
	c.views = append(c.views, v)

	return v, nil
}
