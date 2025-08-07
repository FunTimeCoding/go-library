package constant

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
	assert.String(t, "text/markdown", MarkdownContentType)
	assert.String(t, ":9090", MetricsAddress)
	assert.String(t, "image/x-icon", IconContentType)
	assert.String(t, "Basic", BasicPrefix)
}
