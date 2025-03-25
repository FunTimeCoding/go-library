package age_fixture

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestColorable(t *testing.T) {
	assert.True(t, New(0) != nil)
}
