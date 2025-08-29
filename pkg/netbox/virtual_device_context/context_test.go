package virtual_device_context

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestContext(t *testing.T) {
	assert.NotNil(t, New(&netbox.VirtualDeviceContext{}))
}
