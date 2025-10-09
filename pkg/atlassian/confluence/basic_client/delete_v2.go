package basic_client

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) DeleteV2(path string) string {
	r := web.NewDelete(
		fmt.Sprintf(
			"%s://%s%s%s",
			c.scheme,
			c.host,
			constant.PathPrefix,
			path,
		),
	)
	r.SetBasicAuth(c.user, c.token)

	return web.ReadString(web.Send(web.Client(true), r))
}
