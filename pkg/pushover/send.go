package pushover

import (
	"bytes"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/pushover/constant"
	"github.com/funtimecoding/go-library/pkg/pushover/notification"
	web "github.com/funtimecoding/go-library/pkg/web/constant"
	"github.com/funtimecoding/go-library/pkg/web/locator"
	"net/http"
)

func (c *Client) Send(s string) *http.Response {
	result, e := http.Post(
		locator.New(
			constant.Host,
		).Base(constant.Base).Path(constant.Message).String(),
		web.Object,
		bytes.NewBuffer(
			[]byte(
				notation.Encode(
					notification.New(c.user, c.token, s),
					false,
				),
			),
		),
	)
	errors.PanicOnError(e)

	return result
}
