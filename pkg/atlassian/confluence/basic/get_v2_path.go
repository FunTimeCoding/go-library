package basic

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/go-library/pkg/web/locator"
)

func (c *Client) GetV2Path(p string) string {
	l := locator.New(c.host).Base(constant.Base).Path(p).String()

	return c.GetV2(l)
}
