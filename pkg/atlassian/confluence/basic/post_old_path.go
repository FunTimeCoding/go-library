package basic

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/go-library/pkg/web"
	webConstant "github.com/funtimecoding/go-library/pkg/web/constant"
	"github.com/funtimecoding/go-library/pkg/web/locator"
)

func (c *Client) PostOldPath(
	p string,
	body string,
) string {
	r := web.NewPost(
		locator.New(c.host).Base(constant.OldBase).Path(p).String(),
		body,
	)
	r.SetBasicAuth(c.user, c.token)
	r.Header[webConstant.Accept] = []string{
		webConstant.Object,
	}
	r.Header[webConstant.ContentType] = []string{
		webConstant.Object,
	}

	return web.ReadString(web.Send(web.Client(), r))
}
