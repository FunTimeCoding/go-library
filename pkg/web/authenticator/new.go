package authenticator

import (
	"github.com/funtimecoding/go-library/pkg/web/session_store"
	"os"
)

func New(address []string) *Authenticator {
	if len(address) == 0 {
		if a := os.Getenv(LoginAddressEnvironment); a != "" {
			address = append(address, a)
		}
	}

	return &Authenticator{store: session_store.New(), address: address}
}
