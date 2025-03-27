package content

import "github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"

func (c *Content) formatSpace() string {
	if c.Space != "" {
		return c.Space
	}

	return constant.NoSpace
}
