package authenticator

import "github.com/funtimecoding/go-library/pkg/web/session_store"

type Authenticator struct {
	store   *session_store.Store
	address string
}
