package keepass

import "github.com/funtimecoding/go-library/pkg/system/environment"

func NewEnvironment() *Client {
	return New(
		environment.Get(DatabaseEnvironment),
		environment.Get(PasswordEnvironment),
	)
}
