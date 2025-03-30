package confluence

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic_client/response"
	"github.com/funtimecoding/go-library/pkg/notation"
)

func (c *Client) SpacesBasic() []*response.Space {
	var result *response.Spaces
	notation.DecodeStrict(
		c.basic.GetV2("/spaces"),
		&result,
		false,
	)

	return result.Results
}
