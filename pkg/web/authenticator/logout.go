package authenticator

import (
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/funtimecoding/go-library/pkg/web/request_context"
)

func (a *Authenticator) Logout(c *request_context.Context) {
	if s := c.Cookie(web.SessionCookie); s != nil {
		a.store.Remove(s.Value)
		c.UnsetCookie(web.SessionCookie)
	}
}
