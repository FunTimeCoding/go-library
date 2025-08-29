package user_map

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestMap(t *testing.T) {
	assert.NotNil(t, New(nil, []string{}))
}
