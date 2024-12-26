package authenticator

import (
	"github.com/funtimecoding/go-library/pkg/network"
	"github.com/funtimecoding/go-library/pkg/web/session_store"
	"os"
)

func New() *Authenticator {
	var login string

	if a := os.Getenv(LoginAddressEnvironment); a != "" {
		login = a
	} else {
		login = network.LocalhostAddressString
	}

	return &Authenticator{store: session_store.New(), loginAddress: login}
}
