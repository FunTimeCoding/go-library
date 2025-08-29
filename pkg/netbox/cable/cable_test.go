package cable

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestCable(t *testing.T) {
	assert.NotNil(t, New(&netbox.Cable{}))
}
