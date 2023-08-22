package web

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "User-Agent", UserAgentHeader)
	assert.String(t, "http", InsecureScheme)
	assert.String(t, "https", SecureScheme)
	assert.String(t, "Authorization", AuthorizationHeader)
	assert.String(t, "Accept", AcceptHeader)
}
