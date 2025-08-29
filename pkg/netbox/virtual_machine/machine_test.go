package virtual_machine

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestMachine(t *testing.T) {
	assert.NotNil(t, New(&netbox.VirtualMachineWithConfigContext{}))
}
