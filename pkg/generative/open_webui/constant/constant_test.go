package constant

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "OPEN_WEBUI_HOST", HostEnvironment)
	assert.String(t, "OPEN_WEBUI_TOKEN", TokenEnvironment)
}
