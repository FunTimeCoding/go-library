package confluence

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/content"
)

func (c *Client) LabeledKaos() []*content.Content {
	var result []*content.Content

	for _, label := range c.labels {
		result = append(
			result,
			c.SearchKaos(fmt.Sprintf("label=\"%s\"", label))...,
		)
	}

	return result
}
