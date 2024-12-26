package authenticator

import "github.com/funtimecoding/go-library/pkg/web/session_store"

func New(address []string) *Authenticator {
	return &Authenticator{
		store:   session_store.New(),
		address: address,
	}
}
