package user_map

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestMap(t *testing.T) {
	assert.True(t, New(nil, []string{}) != nil)
}
