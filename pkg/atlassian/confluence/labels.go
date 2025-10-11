package confluence

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/go-library/pkg/notation"
)

func (c *Client) Labels() []*response.LabelResult {
	var r *response.Labels
	notation.DecodeStrict(c.basic.GetV2("/labels"), &r, false)

	return r.Results
}
