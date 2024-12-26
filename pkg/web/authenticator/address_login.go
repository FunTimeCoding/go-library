package authenticator

import (
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/funtimecoding/go-library/pkg/web/request_context"
)

func (a *Authenticator) AddressLogin(c *request_context.Context) string {
	address := c.Address()

	if a.address != address {
		return ""
	}

	s := c.Cookie(web.SessionCookie)

	if s == nil {
		identifier := a.store.Create(address)
		s = c.SetCookie(web.SessionCookie, identifier)
	}

	return s.Value
}
