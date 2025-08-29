package age_fixture

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestColorable(t *testing.T) {
	assert.NotNil(t, New(0))
}
