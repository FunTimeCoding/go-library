package power_feed

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestFeed(t *testing.T) {
	assert.NotNil(t, New(&netbox.PowerFeed{}))
}
