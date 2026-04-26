package system

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestAddress(t *testing.T) {
	assert.NotNil(t, Address())
}
