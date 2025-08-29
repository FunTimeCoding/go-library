package inventory_item

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestRole(t *testing.T) {
	assert.NotNil(t, New(&netbox.InventoryItem{}))
}
