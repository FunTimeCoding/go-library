package basic

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/prometheus/loki/basic/constant"
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/funtimecoding/go-library/pkg/web/locator"
)

func (c *Client) Get(path string) string {
	r := web.NewGet(
		locator.New(c.host).Base(constant.Base).Path(path).String(),
	)
	r.SetBasicAuth(c.user, c.password)
	response := web.Send(web.Client(true), r)

	if false {
		fmt.Println(r)
		fmt.Println(response)
	}

	return web.ReadString(response)
}
