package basic

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/constant"
	"github.com/funtimecoding/go-library/pkg/strings/join/key_value"
	"github.com/funtimecoding/go-library/pkg/web"
	webConstant "github.com/funtimecoding/go-library/pkg/web/constant"
	"github.com/funtimecoding/go-library/pkg/web/locator"
)

func (c *Client) GetBytes(
	path string,
	query map[string]string,
) []byte {
	l := locator.New(c.host).Base(constant.Base).Path(path).Trail()

	for k, v := range query {
		l.Add(k, v)
	}

	r := web.NewGet(l.String())
	r.Header.Add(
		webConstant.Authorization,
		key_value.Space(webConstant.Bearer, c.token),
	)
	s := web.Send(web.Client(), r)
	defer errors.PanicClose(s.Body)

	return web.ReadBytes(s)
}
