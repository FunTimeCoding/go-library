package network

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestInterface(t *testing.T) {
	assert.NotNil(t, New(&netbox.Interface{}))
}
