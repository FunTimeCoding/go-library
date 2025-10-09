package basic_client

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/go-library/pkg/web"
	webConstant "github.com/funtimecoding/go-library/pkg/web/constant"
)

func (c *Client) PostV2(
	path string,
	body string,
) string {
	r := web.NewPost(
		fmt.Sprintf(
			"%s://%s%s%s",
			c.scheme,
			c.host,
			constant.PathPrefix,
			path,
		),
		body,
	)
	r.SetBasicAuth(c.user, c.token)
	r.Header[webConstant.AcceptHeader] = []string{
		webConstant.ObjectContentType,
	}
	r.Header[webConstant.ContentTypeHeader] = []string{
		webConstant.ObjectContentType,
	}

	return web.ReadString(web.Send(web.Client(true), r))
}
