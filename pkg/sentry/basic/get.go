package basic

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/sentry/constant"
	"github.com/funtimecoding/go-library/pkg/strings/join/key_value"
	"github.com/funtimecoding/go-library/pkg/web"
	webConstant "github.com/funtimecoding/go-library/pkg/web/constant"
	"github.com/funtimecoding/go-library/pkg/web/locator"
)

func (c *Client) Get(path string) string {
	r := web.NewGet(
		locator.New(c.host).Base(constant.Base).Path(path).String(),
	)
	r.Header.Add(webConstant.ContentType, webConstant.Object)
	r.Header.Add(
		webConstant.Authorization,
		key_value.Space(webConstant.Bearer, c.token),
	)
	r.Header.Add(webConstant.Accept, webConstant.Object)
	response := web.Send(web.Client(true), r)
	fmt.Println(r)
	fmt.Println(response)

	return web.ReadString(response)
}
