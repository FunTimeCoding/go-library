package habitica

import "github.com/funtimecoding/go-library/pkg/system/environment"

func NewEnvironment() *Client {
	return New(
		environment.Required(HostEnvironment),
		environment.Required(UserEnvironment),
		environment.Required(TokenEnvironment),
	)
}
