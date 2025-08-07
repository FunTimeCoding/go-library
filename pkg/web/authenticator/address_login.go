package authenticator

import (
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"github.com/funtimecoding/go-library/pkg/web/request_context"
)

func (a *Authenticator) AddressLogin(c *request_context.Context) string {
	address := c.Address()

	if address != a.loginAddress {
		return ""
	}

	s := c.Cookie(constant.SessionCookie)

	if s == nil {
		identifier := a.store.Create(address)
		s = c.SetCookie(constant.SessionCookie, identifier)
	}

	return s.Value
}
