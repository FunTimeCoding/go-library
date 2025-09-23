package discord

import "github.com/funtimecoding/go-library/pkg/system/environment"

func NewEnvironment() *Client {
	return New(environment.Required(TokenEnvironment))
}
