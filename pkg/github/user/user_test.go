package user

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestUser(t *testing.T) {
	assert.True(t, New(nil) != nil)
}
