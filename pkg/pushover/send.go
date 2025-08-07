package pushover

import (
	"bytes"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/pushover/constant"
	"github.com/funtimecoding/go-library/pkg/pushover/notification"
	web "github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func (c *Client) Send(s string) *http.Response {
	result, e := http.Post(
		constant.Locator,
		web.ObjectContentType,
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
