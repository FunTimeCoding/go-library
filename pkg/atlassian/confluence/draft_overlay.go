package confluence

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/web/locator"
)

func (c *Client) DraftOverlay(identifier string) (*page.Page, error) {
	body, e := c.basic.Get(
		locator.New(c.host).Base(constant.OldBase).Path(
			"/content/%s",
			identifier,
		).Set(constant.Status, constant.DraftStatus).
			Set("expand", "body.storage,version").String(),
	)

	if e != nil {
		return nil, e
	}

	var result *response.Page
	notation.MustDecode(body, &result, false)

	return page.New(result, c.host), nil
}
