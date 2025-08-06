package user

import (
	"testing"

	"github.com/funtimecoding/go-library/pkg/assert"
)

func TestUser(t *testing.T) {
	assert.True(t, New() != nil)
}
