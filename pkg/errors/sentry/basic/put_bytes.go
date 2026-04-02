package basic

import (
	"bytes"
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/constant"
	"github.com/funtimecoding/go-library/pkg/strings/join/key_value"
	"github.com/funtimecoding/go-library/pkg/web"
	webConstant "github.com/funtimecoding/go-library/pkg/web/constant"
	"github.com/funtimecoding/go-library/pkg/web/locator"
)

func (c *Client) PutBytes(path string, body any) []byte {
	b, e := json.Marshal(body)
	errors.PanicOnError(e)
	l := locator.New(c.host).Base(constant.Base).Path(path).Trail().String()
	r := web.NewPutBytes(l, bytes.NewReader(b))
	r.Header.Add(
		webConstant.Authorization,
		key_value.Space(webConstant.Bearer, c.token),
	)
	r.Header.Add(webConstant.ContentType, webConstant.Object)
	r.Header.Add(webConstant.Accept, webConstant.Object)
	s := web.Send(web.Client(), r)
	defer errors.PanicClose(s.Body)

	return web.ReadBytes(s)
}
