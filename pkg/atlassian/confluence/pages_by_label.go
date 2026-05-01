package confluence

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/go-library/pkg/notation"
)

func (c *Client) PagesByLabel(labelIdentifier string) ([]*page.Page, error) {
	body, e := c.basic.GetV2(
		c.basic.Base().Copy().Path(
			"%s/%s%s",
			constant.Label,
			labelIdentifier,
			constant.Page,
		).Set(constant.BodyFormat, constant.StorageFormat).String(),
	)

	if e != nil {
		return nil, e
	}

	var result *response.Pages
	notation.DecodeStrict(body, &result, false)

	return page.NewSlice(result.Results, c.host), nil
}
