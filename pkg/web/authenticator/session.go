package authenticator

import (
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/funtimecoding/go-library/pkg/web/request_context"
)

func (a *Authenticator) Session(c *request_context.Context) string {
	if s := c.Cookie(web.SessionCookie); s != nil {
		return s.Value
	}

	return ""
}
