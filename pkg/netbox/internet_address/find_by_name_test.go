package internet_address

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

const fixtureAddress = "192.168.0.1/24"

func TestFindByName(t *testing.T) {
	addresses := []*Address{{Name: fixtureAddress}}

	// Happy path
	assert.Any(
		t,
		&Address{Name: "192.168.0.1/24"},
		FindByName(addresses, "192.168.0.1/24"),
	)

	// Not found
	var expected *Address
	assert.Any(t, expected, FindByName(addresses, "192.168.0.2/24"))
}
