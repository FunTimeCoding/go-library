package user_map

import (
	"testing"

	"github.com/funtimecoding/go-library/pkg/assert"
)

func TestMap(t *testing.T) {
	assert.True(t, New(nil, []string{}) != nil)
}
