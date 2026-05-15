package mock_client

import "github.com/funtimecoding/go-library/pkg/sublime/view"

func (c *Client) CreateView(
	title string,
	content string,
	syntax string,
) (view.View, error) {
	v := view.View{
		Identifier: c.nextID,
		Title:      title,
		Text:       content,
		Syntax:     syntax,
	}
	c.nextID++
	c.views = append(c.views, v)

	return v, nil
}
