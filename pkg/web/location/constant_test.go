package location

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "/favicon.ico", Favicon)
	assert.String(t, "/shutdown", Shutdown)
	assert.String(t, "/status", Status)
}
