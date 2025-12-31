package constant

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(
		t,
		"/.well-known/oauth-protected-resource",
		ProtectedResource,
	)
	assert.String(t, "authorization_servers", AuthorizationServer)
	assert.String(t, "resource", Resource)
	assert.String(t, "GET, OPTIONS", ProtectedMethods)
}
