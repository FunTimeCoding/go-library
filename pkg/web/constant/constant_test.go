package constant

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "*", OriginAll)
	assert.String(t, "8080", ListenPort)
	assert.String(t, "Access-Control-Allow-Headers", AccessHeader)
	assert.String(t, "Access-Control-Allow-Methods", AccessMethod)
	assert.String(t, "Access-Control-Allow-Origin", AccessOrigin)
	assert.String(t, "Access-Control-Expose-Headers", AccessExpose)
	assert.String(t, "User-Agent", UserAgent)
	assert.String(t, "WWW-Authenticate", Authenticate)
	assert.String(t, "image/x-icon", Icon)
	assert.String(t, "text/markdown", Markdown)
}
