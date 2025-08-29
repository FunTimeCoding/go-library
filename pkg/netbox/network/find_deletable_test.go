package network

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestFindDeletable(t *testing.T) {
	assert.Any(
		t,
		[]*Interface{{Name: "eth1"}},
		FindDeletable([]*Interface{{Name: Eth1}}, []*Definition{{Name: Eth0}}),
	)
}
