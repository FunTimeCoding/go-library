package authenticator

import "github.com/funtimecoding/go-library/pkg/web/request_context"

func (a *Authenticator) LoggedIn(c *request_context.Context) bool {
	return a.Session(c) != ""
}
