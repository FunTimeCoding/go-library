package confluence

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic_client/response"
	"github.com/funtimecoding/go-library/pkg/notation"
)

func (c *Client) Labels() []*response.LabelResult {
	var result *response.Labels
	notation.DecodeStrict(
		c.basic.GetV2("/labels"),
		&result,
		false,
	)

	return result.Results
}
