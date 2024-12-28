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
	assert.String(t, "multipart/form-data", FormDataContentType)
	assert.String(t, ":9090", MetricsAddress)
}
