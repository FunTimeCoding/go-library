package system

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestCleanAddress(t *testing.T) {
	assert.String(t, "0.0.0.0", CleanAddressStrict("0.0.0.0:22"))
	assert.String(t, "0.0.0.0", CleanAddressStrict("0.0.0.0/32"))
}
