package system

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestUser(t *testing.T) {
	assert.NotNil(t, User())
}
