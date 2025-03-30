package confluence

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic_client/response"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/go-library/pkg/notation"
)

func (c *Client) PageBasic(identifier string) *response.Page {
	var result *response.Page
	notation.DecodeStrict(
		c.basic.GetV2(
			fmt.Sprintf(
				"/pages/%s?body-format=%s",
				identifier,
				constant.StorageFormat,
			),
		),
		&result,
		false,
	)

	return result
}
