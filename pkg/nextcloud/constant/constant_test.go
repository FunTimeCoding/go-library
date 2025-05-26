package constant

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "NEXTCLOUD_HOST", HostEnvironment)
	assert.String(t, "NEXTCLOUD_USER", UserEnvironment)
	assert.String(t, "NEXTCLOUD_PASSWORD", PasswordEnvironment)
}
