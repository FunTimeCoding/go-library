package confluence

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/content"
)

func (c *Client) Labeled() []*content.Content {
	var result []*content.Content

	for _, label := range c.labels {
		result = append(
			result,
			c.Search(fmt.Sprintf("label=\"%s\"", label))...,
		)
	}

	return result
}
