package basic

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/go-library/pkg/web/locator"
)

func (c *Client) GetV2Path(p string) (string, error) {
	return c.GetV2(
		locator.New(c.host).Base(constant.Base).Path(p).String(),
	)
}
