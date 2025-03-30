package confluence

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic_client/response"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/go-library/pkg/notation"
)

func (c *Client) PagesBasic() []*response.Page {
	var result *response.Pages
	notation.DecodeStrict(
		c.basic.GetV2(
			fmt.Sprintf(
				"/pages?body-format=%s",
				constant.StorageFormat,
			),
		),
		&result,
		false,
	)

	return result.Results
}
