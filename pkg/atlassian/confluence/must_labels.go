package confluence

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) MustLabels() []*response.LabelResult {
	result, e := c.Labels()
	errors.PanicOnError(e)

	return result
}
