package basic

import (
	"bytes"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/go-library/pkg/web"
	webConstant "github.com/funtimecoding/go-library/pkg/web/constant"
	"github.com/funtimecoding/go-library/pkg/web/locator"
)

func (c *Client) PutV2Path(
	p string,
	body string,
) string {
	r := web.NewPutBytes(
		locator.New(c.host).Base(constant.Base).Path(p).String(),
		bytes.NewBuffer([]byte(body)),
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
