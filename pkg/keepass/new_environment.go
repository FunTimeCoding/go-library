package keepass

import "github.com/funtimecoding/go-library/pkg/system/environment"

func NewEnvironment() *Client {
	return New(
		environment.Required(DatabaseEnvironment),
		environment.Required(PasswordEnvironment),
	)
}
