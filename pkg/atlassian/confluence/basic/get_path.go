package basic

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/go-library/pkg/web/locator"
)

func (c *Client) GetPath(path string) string {
	return c.Get(
		locator.New(c.host).Base(constant.OldBase).Path(path).String(),
	)
}
