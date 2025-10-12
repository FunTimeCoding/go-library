package basic

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/go-library/pkg/web/locator"
)

func (c *Client) DeleteV2Path(p string) string {
	return c.DeleteV2(
		locator.New(c.host).Base(constant.Base).Path(p).String(),
	)
}
