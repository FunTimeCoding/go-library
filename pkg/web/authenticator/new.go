package authenticator

import (
	"github.com/funtimecoding/go-library/pkg/network"
	"github.com/funtimecoding/go-library/pkg/web/session_store"
	"os"
)

func New() *Authenticator {
	var address []string

	if a := os.Getenv(LoginAddressEnvironment); a != "" {
		address = append(address, a)
	} else {
		address = append(address, network.LocalhostAddressString)
	}

	return &Authenticator{store: session_store.New(), address: address}
}
