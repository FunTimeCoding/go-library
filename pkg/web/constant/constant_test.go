package constant

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "8080", ListenPort)
	assert.String(t, "Accept-Language", AcceptLanguage)
	assert.String(t, "User-Agent", UserAgent)
	assert.String(t, "image/x-icon", Icon)
}
