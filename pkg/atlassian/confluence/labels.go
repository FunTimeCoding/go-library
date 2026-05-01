package confluence

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/go-library/pkg/notation"
)

func (c *Client) Labels() ([]*response.LabelResult, error) {
	body, e := c.basic.GetV2Path(constant.Label)

	if e != nil {
		return nil, e
	}

	var r *response.Labels
	notation.DecodeStrict(body, &r, false)

	return r.Results, nil
}
