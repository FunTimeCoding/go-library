package confluence

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/page"
	"slices"
)

func (c *Client) Labeled() []*page.Page {
	var result []*page.Page
	var identifiers []string

	for _, l := range c.Labels() {
		if slices.Contains(c.labels, l.Name) {
			identifiers = append(identifiers, l.Id)
		}
	}

	for _, identifier := range identifiers {
		result = append(result, c.PagesByLabel(identifier)...)
	}

	return result
}
