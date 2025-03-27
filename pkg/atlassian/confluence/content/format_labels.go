package content

import "github.com/funtimecoding/go-library/pkg/strings/join"

func (c *Content) formatLabels() string {
	if len(c.Labels) > 0 {
		return join.CommaSpace(c.Labels)
	}

	return ""
}
